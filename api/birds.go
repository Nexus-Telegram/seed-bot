package api

import (
	"errors"
	"github.com/nexus-telegram/seed-bot/types"
	"go.uber.org/zap"
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
func (s *Service) GetBirds() (*types.BirdsData, error) {
	var result types.BirdsData
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

type FeedBirdPayload struct {
	BirdId  string   `json:"bird_id"`
	WormIds []string `json:"worm_ids"`
}

func (s *Service) FeedBird(wormIDs []string, birdID string) (*types.BirdData, error) {
	var BirdFed types.BirdData
	payload := FeedBirdPayload{
		WormIds: wormIDs,
		BirdId:  birdID,
	}
	res, err := s.Client.R().SetBody(payload).SetResult(&BirdFed).Post("/bird-feed")
	if err != nil {
		s.Logger.Error("Error while feeding bird", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error feeding bird",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}
	return &BirdFed, nil

}

type StartBirdHuntPayload struct {
	BirdId    string `json:"bird_id"`
	TaskLevel int    `json:"task_level"`
}

func (s *Service) StartBirdHunt(birdId string, taskLevel int) (*types.BirdData, error) {
	var payload = StartBirdHuntPayload{
		BirdId:    birdId,
		TaskLevel: taskLevel,
	}
	var startedBird types.BirdData
	res, err := s.Client.R().SetBody(payload).SetResult(startedBird).Post("/bird-hunt/start")
	if err != nil {
		s.Logger.Error("Error while starting bird hunt", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Error starting hunt",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}

	return &startedBird, nil
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
