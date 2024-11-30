package api

import (
	"github.com/nexus-telegram/seed-bot/types"
	"net/http"

	"go.uber.org/zap"
)

func (s *Service) GetLoginBonus() *types.LoginBonuses {
	var loginBonuses types.LoginBonuses
	res, err := s.Client.R().SetResult(&loginBonuses).Get("/login-bonuses")
	if err != nil {
		s.Logger.Error("Error while fetching login bonuses", zap.Error(err))
		return nil
	}
	if res.IsError() {
		s.Logger.Error("Failed to fetch login bonuses",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil
	}
	return &loginBonuses
}

func (s *Service) GetStreakReward() *types.StreakRewards {
	var streakRewards types.StreakRewards
	res, err := s.Client.R().SetResult(&streakRewards).Get("/streak-reward")
	if err != nil {
		s.Logger.Error("Error while fetching streak rewards", zap.Error(err))
		return nil
	}
	if res.IsError() {
		s.Logger.Error("Failed to claim streak rewards",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil
	}
	return &streakRewards
}

func (s *Service) GetDailyLoginStreak() *types.LoginStreak {
	var loginStreak types.LoginStreak
	res, err := s.Client.R().SetResult(&loginStreak).Get("/daily-login-streak")
	if err != nil {
		s.Logger.Error("Error while fetching streak rewards", zap.Error(err))
		return nil
	}
	if res.IsError() {
		s.Logger.Error("Failed to fetch daily login streak",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil
	}
	return &loginStreak
}

func (s *Service) ClaimStreakReward(streakRewardIds []string) *types.StreakRewards {
	payload := types.StreakRewardsPayload{
		Ids: streakRewardIds,
	}
	var streakRewards types.StreakRewards
	res, err := s.Client.R().SetBody(payload).SetResult(&streakRewards).Post("/streak-reward")
	if err != nil {
		s.Logger.Error("Error while claiming streak rewards", zap.Error(err))
		return nil
	}
	if res.IsError() {
		s.Logger.Error("Failed to claim streak rewards",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil
	}
	return &streakRewards
}

func (s *Service) ClaimLoginBonus() *types.LoginBonusesCreate {
	var loginBonuses types.LoginBonusesCreate
	res, err := s.Client.R().SetResult(&loginBonuses).Post("/login-bonuses")
	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}
	if res.StatusCode() == http.StatusBadRequest {
		// Verifica se o resultado pode ser convertido para o tipo esperado
		errorResult, ok := res.Result().(*types.ErrorResponse)
		if !ok {
			s.Logger.Error("Unexpected error response type",
				zap.String("status", res.Status()),
				zap.Any("rawBody", string(res.Body())),
			)
			return nil
		}

		// Valida os campos do ErrorResponse
		if errorResult.Code == "invalid-request" && errorResult.Message == "already claimed for today" {
			s.Logger.Info("Daily already claimed")
			return nil
		}

		s.Logger.Error("Failed to fetch login bonuses",
			zap.String("status", res.Status()),
			zap.Any("errorData", errorResult),
		)
		return nil
	}
	if res.IsError() {
		s.Logger.Error("Failed to fetch login bonuses",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil

	}
	return &loginBonuses
}
