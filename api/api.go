package api

import (
	"encoding/json"
	"fmt"
	"log"
	"nexus-seed-bot/types"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Service struct {
	Client    *resty.Client
	Logger    *zap.Logger
	BalanceCh chan int // Channel to send balance updates
}

func NewService(client *resty.Client, logger *zap.Logger) *Service {
	return &Service{
		Client:    client,
		Logger:    logger,
		BalanceCh: make(chan int), // Initialize the channel
	}
}

func (s *Service) GetProfile() (*types.ProfileResponse, error) {
	resp, err := s.Client.R().
		SetResult(&types.ProfileResponse{}).
		Get("/profile2")
	if err != nil {
		return nil, fmt.Errorf("error fetching profile data: %v", err)
	}
	data := resp.Result().(*types.ProfileResponse)
	return data, nil
}

func (s *Service) GetSeed() error {
	resp, err := s.Client.R().
		Post("/seed/claim")
	if err != nil {
		return fmt.Errorf("error claiming seed: %v", err)
	}

	if resp.StatusCode() == 200 {
		s.Logger.Info("✅ SEED claimed successfully!")
	}
	return nil
}

func (s *Service) GetNextWormTime() *types.CatchMetadataResponse {
	wormsRes, err := s.Client.R().SetResult(&types.CatchMetadataResponse{}).Get("/worms")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return wormsRes.Result().(*types.CatchMetadataResponse)
}

func (s *Service) GetWorm() {
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
func (s *Service) CompleteTask(taskID string) error {
	payload := "{}"
	res, err := s.Client.R().SetBody(payload).Post(fmt.Sprintf("/tasks/%s", taskID))
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	if res.StatusCode() == 200 {
		s.Logger.Info(fmt.Sprintf("Success completing task: %s", taskID))
	}
	return nil
}
func (s *Service) GetProgress() *types.Progress {
	res, err := s.Client.R().SetResult(&types.Progress{}).Get("/tasks/progresses")
	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}
	return res.Result().(*types.Progress)
}
func (s *Service) GetBalance() int {
	response, err := s.Client.R().SetResult(&types.Balance{}).Get("/profile/balance")
	if err != nil {
		s.Logger.Error("Error while fetching balance:", zap.Error(err))
	}
	balance := response.Result().(*types.Balance).Balance

	return balance
}

func (s *Service) GetSettings() (*types.Settings, error) {
	response, err := s.Client.R().SetResult(&types.Settings{}).Get("/settings")
	if err != nil {
		s.Logger.Error("Error while fetching settings:", zap.Error(err))
		return nil, err
	}
	settings := response.Result().(*types.Settings)
	return settings, nil
}
func (s *Service) BuyUpgrade() error {
	_, err := s.Client.R().Post("/seed/mining-speed/upgrade")
	if err != nil {
		s.Logger.Error("Error while buying upgrade:", zap.Error(err))
		return err
	}
	return nil
}
