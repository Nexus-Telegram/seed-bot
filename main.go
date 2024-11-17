package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nexus-seed-bot/types"
	"time"

	nexuslogger "github.com/Nexus-Telegram/nexus/logger"

	"github.com/Nexus-Telegram/nexus/utils"
	"go.uber.org/zap"

	"github.com/go-resty/resty/v2"
)

func getProfile(client *resty.Client) (*types.ProfileResponse, error) {
	resp, err := client.R().
		SetResult(&types.ProfileResponse{}).
		Get("/profile2")
	if err != nil {
		return nil, fmt.Errorf("error fetching profile data: %v", err)
	}
	data := resp.Result().(*types.ProfileResponse)
	return data, nil
}

func getSeed(client *resty.Client, logger *zap.Logger) error {
	resp, err := client.R().
		Post("/seed/claim")
	if err != nil {
		return fmt.Errorf("error claiming seed: %v", err)
	}

	if resp.StatusCode() == 200 {
		logger.Info("✅ SEED claimed successfully!")
	}
	return nil
}

func autoCompleteTasks(client *resty.Client) {
	fmt.Println("\n✅ Auto completing tasks...")

	tasksResp, err := client.R().SetResult(&types.TaskResponse{}).
		Get("/tasks/progresses")
	if err != nil {
		log.Fatalf("Error fetching tasks: %v", err)
	}
	tasksRespData := tasksResp.Result().(*types.TaskResponse)

	for i := range tasksRespData.Data {
		fmt.Println(i)
	}
}

func getCommonHeaders() map[string]string {
	return map[string]string{
		"accept":             "*/*",
		"accept-language":    "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7,pt-PT;q=0.6",
		"origin":             "https://cf.seeddao.org",
		"priority":           "u=1, i",
		"referer":            "https://cf.seeddao.org/",
		"sec-ch-ua":          `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Linux"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-site",
		"user-agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	}
}

func getNextWormTime(client *resty.Client, logger *zap.Logger) *types.CatchMetadataResponse {
	wormsRes, err := client.R().SetResult(&types.CatchMetadataResponse{}).Get("/worms")
	if err != nil {
		logger.Error(err.Error())
	}
	return wormsRes.Result().(*types.CatchMetadataResponse)
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func getWorm(client *resty.Client, logger *zap.Logger) {
	res, err := client.R().Post("/worms/catch")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode() == 400 {
		logger.Info("Worm already caught")
		return
	}
	if res.StatusCode() == 404 {
		var errResp ErrorResponse
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			logger.Error("Failed to unmarshal response", zap.Error(err))
			return
		}
		if errResp.Code == "resource-not-found" && errResp.Message == "worm not found" {
			logger.Error(fmt.Sprintf("Error: %s, Message: %s", errResp.Code, errResp.Message))
		}
		return
	}

	if res.StatusCode() == 401 {
		var errResp ErrorResponse
		// SetResult automatically unmarshals the body into the provided struct
		err := json.Unmarshal(res.Body(), &errResp)
		if err != nil {
			logger.Error("Failed to unmarshal response", zap.Error(err))
			return
		}

		if errResp.Code == "authentication" && errResp.Message == "telegram data expired" {
			logger.Error("Telegram data expired")
			// Handle the error as needed, e.g., request a new authentication token
		}
	}
}

func waitUntilNextWorm(nextWormTime time.Time, logger *zap.Logger) {
	durationUntilNextWorm := time.Until(nextWormTime.Add(1 * time.Minute))
	if durationUntilNextWorm > 0 {
		logger.Info(fmt.Sprintf("Waiting until %s to catch the worm.", nextWormTime))
		time.Sleep(durationUntilNextWorm)
		logger.Info("It's time to catch the worm!")
	} else {
		logger.Info("The next worm time has already passed.")
	}
}

func runWormCatching(client *resty.Client, logger *zap.Logger) {
	for {
		wormsMetaData := getNextWormTime(client, logger)

		if wormsMetaData.Data.IsCaught {
			waitUntilNextWorm(wormsMetaData.Data.NextWorm, logger)
		} else {
			getWorm(client, logger)
		}
	}
}
func runSeedClaim(client *resty.Client, logger *zap.Logger) {
	for {
		profile, err := getProfile(client)
		if err != nil {
			logger.Error(err.Error())
			continue // Skip to the next iteration if getProfile fails
		}

		timeUntilClaim := time.Until(profile.Data.LastClaim)

		if timeUntilClaim > 3*time.Hour {
			seedErr := getSeed(client, logger)
			if seedErr != nil {
				logger.Error(seedErr.Error()) // Log seed error properly
			}
		} else {
			// Sleep until the next claim time
			if timeUntilClaim > 0 { // Ensure timeUntilClaim is positive
				time.Sleep(timeUntilClaim)
			}
		}
	}
}
func main() {
	client := resty.New().SetTimeout(30 * time.Second)
	client.SetBaseURL("https://alb.seeddao.org/api/v1")
	logger, err := nexuslogger.NewLogger(false)
	if err != nil {
		panic(err)
	}
	queryIDs := utils.ParseQueryIDs()
	for _, queryID := range queryIDs {
		logger.Info(fmt.Sprintf("%s", queryID))

		authHeaders := map[string]string{
			"Telegram-Data": queryID,
		}
		headers := utils.MergeHeaders(getCommonHeaders(), authHeaders)
		client.SetHeaders(headers)
		go runWormCatching(client, logger)
		go runSeedClaim(client, logger)

	}
	select {}
}
