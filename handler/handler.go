package handler

import (
	"fmt"
	"github.com/nexus-telegram/seed-bot/api"
	"github.com/nexus-telegram/seed-bot/types"
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
			caughtWorm := s.CatchWorm()
			if caughtWorm != nil {
				if caughtWorm.Data.Status == "successful" {
					s.Logger.Info(fmt.Sprintf("Successfully catch worm of type %s and reward %d", caughtWorm.Data.Type, caughtWorm.Data.Reward))
					caughtWorms := []types.CaughtWorm{
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
								Id:        caughtWorm.Data.Id,
								Type:      caughtWorm.Data.Type,
								Status:    caughtWorm.Data.Status,
								UpdatedAt: caughtWorm.Data.UpdatedAt,
								Reward:    caughtWorm.Data.Reward,
								OnMarket:  caughtWorm.Data.OnMarket,
								OwnerId:   caughtWorm.Data.OwnerId,
							},
						},
					}
					s.WormCh <- caughtWorms
				}
			}
		}
	}
}

func HandleTasks(s *api.Service) {
	progresses := s.GetProgress()
	if progresses == nil {
		s.Logger.Error("Error while fetching progress")
		return
	}
	// TODO: Implement upgrade tasks
	//upgradeTasks, err := s.getUpgradeTasks()
	//if err != nil {
	//	s.Logger.Error("Error while fetching upgrade tasks")
	//}
	//for _, upgradeTask := range upgradeTasks {
	//	s.CompleteUpgradeTasks(task.Id)
	//}

	for _, task := range progresses.Data {
		if task.TaskUser != nil {
			continue
		}
		if task.Type == types.FollowUs || task.Type == types.PlayApp || task.Type == types.AddList {
			err := s.CompleteTaskWithoutConfirmation(task.Id)
			if task.Id == "dc1bc321-a395-422a-be0d-f20bb6234d6e" || task.Id == "b380ece7-5b68-4a1e-9344-889769633d14" {
				go func(task types.UniqueTask) {
					time.Sleep(20 * time.Second)
					err := s.CompleteTaskWithoutConfirmation(task.Id)
					if err != nil {
						s.Logger.Error(fmt.Sprintf("Error while completed task %s that need two times", task.Name))
					}
					s.Logger.Info(fmt.Sprintf("Task %s completed successfull", task.Name))
				}(task)
			}
			if err != nil {
				s.Logger.Error(err.Error())
			}
			s.Logger.Info(fmt.Sprintf("Task %s completed successfully ", task.Name))
		}

		if task.Type == types.Academy {
			secret, exists := types.TaskSecrets[task.Id]
			if exists {
				err, _ := s.CompleteTaskWithSecret(task.Id, secret)
				if err != nil {
					s.Logger.Debug(fmt.Sprintf("Task %s completed successfully", task.Name))
					continue
				}
				continue
			}
			s.Logger.Error("Secret Not found for task", zap.Any("task", task))
			continue
		}
		continue
	}
}

//	func HandleUpgrade(s *api.Service) {
//		for {
//			balance := <-s.BalanceCh
//			s.Logger.Info(fmt.Sprintf("Balance updated %d", balance))
//
//			profile, err := s.GetProfile()
//			if err != nil {
//				s.Logger.Error("Error fetching profile", zap.Error(err))
//				continue
//			}
//
//			upgrades := profile.Data.Upgrades
//			var upgradeLevel int
//			if len(upgrades) > 0 {
//				upgradeLevel = upgrades[0].UpgradeLevel
//			}
//			settings, err := s.GetSettings()
//			if err != nil {
//				s.Logger.Error(err.Error())
//				continue
//			}
//			UpgradePrices := settings.Data.MiningSpeed
//			if upgradeLevel < len(UpgradePrices) {
//				requiredBalance := UpgradePrices[upgradeLevel]
//
//				if balance >= requiredBalance {
//					err := s.BuyUpgrade()
//					if err != nil {
//						s.Logger.Error("Error while buying upgrade", zap.Error(err))
//					} else {
//						s.Logger.Info("Upgrade purchased successfully", zap.Int("level", upgradeLevel+1))
//					}
//				} else {
//					s.Logger.Info("Insufficient balance for next upgrade", zap.Int("required", requiredBalance), zap.Int("current", balance))
//				}
//			} else {
//				s.Logger.Info("Max upgrade level reached")
//			}
//		}
//	}
func HandleInitializeBird(s *api.Service) {
	myEggs := s.GetMyEggs()
	if myEggs.Total > 0 {
		return
	}
	birds, err := s.GetMyBirds()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}
	if birds.Data.Total > 0 {
		s.Logger.Info("BirdsData already initialized, skipping...")
		return
	}
	firstEgg, err := s.TakeFirstEgg()
	if err != nil {
		s.Logger.Error("error while catching first egg", zap.Error(err))
		return
	}
	hatchEggErr := s.HatchEgg(firstEgg.Id)
	if hatchEggErr != nil {
		s.Logger.Error("error while catching hatch egg", zap.Error(hatchEggErr))
		return
	}
	s.Logger.Info("all initialization tasks successfully")
}

func birdIsReadyToComplete(bird types.Bird) bool {
	if bird.Status == "hunting" && bird.HuntEndAt.Before(time.Now().UTC()) {
		return true
	} else {
		return false
	}
}

func HandleBird(s *api.Service) {
	birds, err := s.GetBirds()
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}
	for _, bird := range birds.Bird {
		// TODO: Get the bird to check if it can be clicked before making the POST request
		_, err := s.ClickBird(bird.Id)
		if err != nil {
			s.Logger.Error(err.Error())
		}
		var wormIds []string
		caughtWorm := <-s.WormCh
		for _, worm := range caughtWorm {
			wormIds = append(wormIds, worm.Data.Id)
		}
		_, err = s.FeedBird(wormIds, bird.Id)
		if err != nil {
			s.Logger.Error(err.Error())
		}
	}
	// TODO: Wait the birdHunting channel when a bird gets in Hunting status then wait the 12 hours for Complete it
}

func isInList(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func WaitUntilNewDay() {
	now := time.Now().UTC()

	nextDay := now.Add(24 * time.Hour).Truncate(24 * time.Hour)

	nextDayWithOneHour := nextDay.Add(10 * time.Minute)

	time.Sleep(time.Until(nextDayWithOneHour))
}

func HandleDaily(s *api.Service) {
	loginStreak := s.GetDailyLoginStreak()
	streakRewards := s.GetStreakReward()
	for _, reward := range streakRewards.Data {
		var createdRewards []string
		if reward.Status == types.Created {
			createdRewards = append(createdRewards, reward.Id)
		}
		if len(createdRewards) > 0 {
			s.ClaimStreakReward(createdRewards)
			s.Logger.Info("Rewards claimed successfully")
		}
	}

	if loginStreak.Data.No == 0 || loginStreak.Data.CreatedAt.Day() != time.Now().UTC().Day() {
		loginBonus := s.ClaimLoginBonus()
		if loginBonus.Data.Timestamp.Day() != time.Now().UTC().Day() {
			s.Logger.Error("Daily Quest failed to be claimed", zap.Any("loginBonus", loginBonus))
			return
		}

		streakReward := s.GetStreakReward()
		for _, oneReward := range streakReward.Data {
			if oneReward.Status == types.Created {
				var rewardCollector []string
				for _, reward := range streakReward.Data {
					rewardCollector = append(rewardCollector, reward.Id)
				}

				streakRewards := s.ClaimStreakReward(rewardCollector)
				for _, claimedStreakReward := range streakRewards.Data {
					if claimedStreakReward.Status == types.Received && isInList(rewardCollector, claimedStreakReward.Id) {
						s.Logger.Info("Daily quest claimed successfully")
					}
				}

			}
		}
	}
	WaitUntilNewDay()

	// After waiting for the next day, you can call the function again if needed.
	// You can loop back here to continue the process.
	HandleDaily(s)
}
