package handler

import (
	"fmt"
	"nexus-seed-bot/api"
	"time"

	"go.uber.org/zap"
)

func HandleSeedClaim(s *api.Service) {
	for {
		profile, err := s.GetProfile()
		if err != nil {
			s.Logger.Error(err.Error())
			return
		}

		timeUntilLastClaim := time.Until(profile.Data.LastClaim)

		// Subtract 2 hours and 1 minute from timeUntilClaim
		TimeUntilClaim := timeUntilLastClaim + (2*time.Hour + 10*time.Second)
		if TimeUntilClaim <= 0 {
			seedErr := s.GetSeed()
			if seedErr != nil {
				s.Logger.Error(seedErr.Error())
				continue
			}
			s.Logger.Info("Seed claimed")
		}

		s.Logger.Info(fmt.Sprintf("Waiting for %02dh%02dm to claim seed", int(TimeUntilClaim.Hours()), int(TimeUntilClaim.Minutes())%60))
		time.Sleep(TimeUntilClaim)
	}
}

func HandleWormCatching(s *api.Service) {
	for {
		wormsMetaData := s.GetNextWormTime()
		if wormsMetaData.Data.IsCaught {
			nextWormTime := wormsMetaData.Data.NextWorm
			durationUntilNextWorm := time.Until(nextWormTime.Add(10 * time.Second))
			s.Logger.Info(fmt.Sprintf("Waiting for %02dh%02dm to catch the worm.", int(durationUntilNextWorm.Hours()), int(durationUntilNextWorm.Minutes())%60))
			time.Sleep(durationUntilNextWorm)

		} else {
			s.CatchWorm()
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
		balance := <-s.BalanceCh
		s.Logger.Info("Received balance update", zap.Int("balance", balance))

		profile, err := s.GetProfile()
		if err != nil {
			s.Logger.Error("Error fetching profile:", zap.Error(err))
			continue
		}

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
func HandleInitialize(s *api.Service) {
	myEggs := s.GetMyEggs()
	if myEggs.Total > 0 {
		return
	}
	firstEgg, err := s.TakeFirstEgg()
	if err != nil {
		HandleInitialize(s)
	}
	hatchEggErr := s.HatchEgg(firstEgg.Id)
	if hatchEggErr != nil {
		return
	}
	s.Logger.Info("initialize completed successfully")

}

func handleBird(s *api.Service) {
	birds, err := s.GetBirds()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}
	for _, bird := range birds.Data {

		_, err := s.ClickBird(bird.Id)
		if err != nil {
			s.Logger.Error(err.Error())
		}
		s.CatchWorm()
	}

}
