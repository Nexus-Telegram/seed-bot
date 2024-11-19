package api

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"nexus-seed-bot/types"
	"time"
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
	var catchedWormResponse types.CatchedWorm
	res, err := s.Client.R().SetResult(&catchedWormResponse).Post("/worms/catch")
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
	if catchedWormResponse.Data.Status == "successful" {
		s.Logger.Info(fmt.Sprintf("Successfully catch worm of type %s and reward %d", catchedWormResponse.Data.Type, catchedWormResponse.Data.Reward))
		catchedWorms := []types.CatchedWorm{
			{
				Data: struct {
					Id        string    `json:"id"`
					Type      string    `json:"type"`
					Status    string    `json:"status"`
					UpdatedAt time.Time `json:"updated_at"`
					Reward    int       `json:"reward"`
					OnMarket  bool      `json:"on_market"`
					OwnerId   string    `json:"owner_id"`
				}{
					Id:        catchedWormResponse.Data.Id,
					Type:      catchedWormResponse.Data.Type,
					Status:    catchedWormResponse.Data.Status,
					UpdatedAt: catchedWormResponse.Data.UpdatedAt,
					Reward:    catchedWormResponse.Data.Reward,
					OnMarket:  catchedWormResponse.Data.OnMarket,
					OwnerId:   catchedWormResponse.Data.OwnerId,
				},
			},
		}
		s.WormCh <- catchedWorms
		return
	}
}
