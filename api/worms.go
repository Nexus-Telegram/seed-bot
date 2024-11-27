package api

import (
	"encoding/json"
	"fmt"
	"github.com/nexus-telegram/seed-bot/types"
	"go.uber.org/zap"
	"log"
)

func (s *Service) GetNextWormTime() *types.CatchMetadataResponse {
	wormsRes, err := s.Client.R().SetResult(&types.CatchMetadataResponse{}).Get("/worms")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return wormsRes.Result().(*types.CatchMetadataResponse)
}

func (s *Service) CatchWorm() *types.CaughtWorm {
	var caughtWormResponse types.CaughtWorm
	res, err := s.Client.R().SetResult(&caughtWormResponse).Post("/worms/catch")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode() == 400 {
		s.Logger.Info("Worm already caught")
		return nil
	}
	if res.StatusCode() == 404 {
		var errResp types.ErrorResponse
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			s.Logger.Error("Failed to unmarshal response", zap.Error(err))
			return nil
		}
		if errResp.Code == "resource-not-found" && errResp.Message == "worm not found" {
			s.Logger.Error(fmt.Sprintf("Error: %s, Message: %s", errResp.Code, errResp.Message))
		}
		return nil
	}

	if res.StatusCode() == 401 {
		var errResp types.ErrorResponse
		// SetResult automatically unmarshal the body into the provided struct
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			s.Logger.Error("Failed to unmarshal response", zap.Error(err))
			return nil
		}

		if errResp.Code == "authentication" && errResp.Message == "telegram data expired" {
			s.Logger.Error("Telegram data expired")
			// Handle the error as needed, e.g., request a new authentication token
		}
	}
	return &caughtWormResponse
}
