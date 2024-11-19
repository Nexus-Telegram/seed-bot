package api

import (
	"errors"
	"go.uber.org/zap"
	"nexus-seed-bot/types"
	"strconv"
)

func (s *Service) ClickBird(birdID string) (*types.BirdsHappiness, error) {
	var birdHappiness types.BirdsHappiness
	payload := map[string]string{
		"bird_id":        birdID,
		"happiness_rate": strconv.Itoa(10000),
	}
	res, err := s.Client.R().SetBody(payload).SetResult(&birdHappiness).Post("/bird-happiness")
	if err != nil {
		s.Logger.Error("Error while clicking bird", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error while clicking bird",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}
	return &birdHappiness, nil

}
func (s *Service) GetMyBirds() (*types.GetMyBirds, error) {
	var result types.GetMyBirds
	res, err := s.Client.R().SetResult(&result).Get("/bird/me?page=1")
	if err != nil {
		s.Logger.Error("Error while fetching birds", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error while fetching birds",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}

	return &result, nil

}
func (s *Service) GetBirds() (*types.Birds, error) {
	var result types.Birds
	res, err := s.Client.R().SetResult(&result).Get("/bird/me-all")
	if err != nil {
		s.Logger.Error("Error while fetching birds", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error while fetching birds",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}

	return &result, nil
}

func FeedBird() {
	//POST
	// https://alb.seeddao.org/api/v1/bird-feed
	//{"bird_id":"47b5ab67-8040-42f9-93ec-82a19b4da89a","worm_ids":["e6b62093-8478-4c6c-a46e-389c64ef0feb"]}
}
func StartBirdHunt() {
	// POST
	//https://alb.seeddao.org/api/v1/bird-hunt/start
	//{"bird_id":"47b5ab67-8040-42f9-93ec-82a19b4da89a","task_level":0}
}

func (s *Service) CompleteBirdHunt(birdID string) (*types.CompleteBirdHunterResponse, error) {
	var birdHuntResponse types.CompleteBirdHunterResponse

	payload := types.CompleteBirdHuntPayload{
		BirdId: birdID,
	}
	res, err := s.Client.R().SetBody(payload).SetResult(&birdHuntResponse).Post("/bird-hunt/complete")
	if err != nil {
		s.Logger.Error("Error while completing bird hunt", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error completing hunting",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}
	return &birdHuntResponse, nil
}
