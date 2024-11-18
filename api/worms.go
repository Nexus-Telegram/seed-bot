package api

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"nexus-seed-bot/types"
)

func GetWorms() {

	//GET
	// https://alb.seeddao.org/api/v1/worms/me-all

}

func (s *Service) GetNextWormTime() *types.CatchMetadataResponse {
	wormsRes, err := s.Client.R().SetResult(&types.CatchMetadataResponse{}).Get("/worms")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return wormsRes.Result().(*types.CatchMetadataResponse)
}

func (s *Service) CatchWorm() {
	res, err := s.Client.R().Post("/worms/catch")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode() == 400 {
		s.Logger.Info("Worm already caught")
		return
	}
	if res.StatusCode() == 404 {
		var errResp types.ErrorResponse
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			s.Logger.Error("Failed to unmarshal response", zap.Error(err))
			return
		}
		if errResp.Code == "resource-not-found" && errResp.Message == "worm not found" {
			s.Logger.Error(fmt.Sprintf("Error: %s, Message: %s", errResp.Code, errResp.Message))
		}
		return
	}

	if res.StatusCode() == 401 {
		var errResp types.ErrorResponse
		// SetResult automatically unmarshals the body into the provided struct
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			s.Logger.Error("Failed to unmarshal response", zap.Error(err))
			return
		}

		if errResp.Code == "authentication" && errResp.Message == "telegram data expired" {
			s.Logger.Error("Telegram data expired")
			// Handle the error as needed, e.g., request a new authentication token
		}
	}
}
