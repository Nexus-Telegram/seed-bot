package handler

import (
	"nexus-seed-bot/api"
	seedUtils "nexus-seed-bot/utils"
	"time"

	"go.uber.org/zap"
)

func HandleSeedClaim(s *api.Service) {
	for {
		profile, err := s.GetProfile()
		if err != nil {
			s.Logger.Error(err.Error())
			continue // Skip to the next iteration if getProfile fails
		}

		timeUntilClaim := time.Until(profile.Data.LastClaim)

		if timeUntilClaim > 3*time.Hour {
			seedErr := s.GetSeed()
			if seedErr != nil {
				s.Logger.Error(seedErr.Error()) // Log seed error properly
			}
		} else {
			// Sleep until the next claim time
			if timeUntilClaim > 0 { // Ensure timeUntilClaim is positive
				time.Sleep(timeUntilClaim)
			}
		}
	}
}

func HandleWormCatching(s *api.Service) {
	for {
		wormsMetaData := s.GetNextWormTime()

		if wormsMetaData.Data.IsCaught {
			seedUtils.WaitUntilNextWorm(wormsMetaData.Data.NextWorm, s.Logger)
		} else {
			s.GetWorm()
			s.GetBalance()
		}
	}
}
func HandleTasks(s *api.Service) {
	progresses := s.GetProgress()
	if progresses == nil {
		s.Logger.Error("Error while fetching progress")
		return
	}
	for _, task := range progresses.Data {
		err := s.CompleteTask(task.Id)
		if err != nil {
			s.Logger.Error(err.Error())
			continue
		}

	}

}
func HandleUpgrade(s *api.Service) {
	for {
		// Aguardar por uma atualização do saldo
		balance := <-s.BalanceCh
		s.Logger.Info("Received balance update", zap.Int("balance", balance))

		// Obter o perfil do usuário
		profile, err := s.GetProfile()
		if err != nil {
			s.Logger.Error("Error fetching profile:", zap.Error(err))
			continue
		}

		// Verificar o nível de upgrade atual
		upgrades := profile.Data.Upgrades
		var upgradeLevel int
		if len(upgrades) > 0 {
			upgradeLevel = upgrades[0].UpgradeLevel
		}
		settings, err := s.GetSettings()
		if err != nil {
			s.Logger.Error(err.Error())
			continue
		}
		UpgradePrices := settings.Data.MiningSpeed
		if upgradeLevel < len(UpgradePrices) {
			requiredBalance := UpgradePrices[upgradeLevel]

			if balance >= requiredBalance {
				err := s.BuyUpgrade()
				if err != nil {
					s.Logger.Error("Error while buying upgrade", zap.Error(err))
				} else {
					s.Logger.Info("Upgrade purchased successfully", zap.Int("level", upgradeLevel+1))
				}
			} else {
				s.Logger.Info("Insufficient balance for next upgrade", zap.Int("required", requiredBalance), zap.Int("current", balance))
			}
		} else {
			s.Logger.Info("Max upgrade level reached")
		}
	}
}
