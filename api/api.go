package api

import (
	"fmt"
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

func (s *Service) GetBalance() int {
	response, err := s.Client.R().SetResult(&types.Balance{}).Get("/profile/balance")
	if err != nil {
		s.Logger.Error("Error while fetching balance:", zap.Error(err))
	}
	balance := response.Result().(*types.Balance).Balance
	s.BalanceCh <- balance

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