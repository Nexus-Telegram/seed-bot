package api

import (
	"errors"
	"go.uber.org/zap"
	"nexus-seed-bot/types"
)

func (s *Service) GetMyEggs() *types.EggData {
	var result types.GetMyEggs
	_, err := s.Client.R().
		SetResult(&result).
		Get("/egg/me?page=1")
	if err != nil {
		s.Logger.Error("Error while fetching Eggs", zap.Error(err))
		return nil
	}

	return &result.Data
}

func (s *Service) TakeFirstEgg() (*types.FirstEggData, error) {
	var result types.FirstEgg
	res, err := s.Client.R().SetResult(&result).Post("/give-first-egg")
	if err != nil {
		s.Logger.Error("Error while taking first egg", zap.Error(err))
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Failed to take first egg",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(string(res.Body()))
	}
	if result.FirstEggData.Status != "in-inventory" {
		s.Logger.Error("Failed to take first egg", zap.Any("data", result))
		return nil, errors.New(string(res.Body()))
	}
	s.Logger.Debug("First egg caught successfully")
	return &result.FirstEggData, nil
}

func (s *Service) HatchEgg(eggID string) error {
	var hatchedEgg types.HatchEgg
	payload := map[string]string{
		"egg_id": eggID,
	}
	res, err := s.Client.R().
		SetBody(payload).
		SetResult(&hatchedEgg).
		Post("/egg-hatch/complete")

	if err != nil {
		s.Logger.Error("Error while hatching egg", zap.Error(err))
		return errors.New("failed to hatch egg ")
	}
	if res.IsError() {
		s.Logger.Error("Failed to hatch egg",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return errors.New("failed to hatch egg ")
	}
	return nil
}

func getEgg() {
	// https://alb.seeddao.org/api/v1/egg/2841614e-6501-4a91-8809-0d8581505010

}
