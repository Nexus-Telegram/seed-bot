package main

import (
	"fmt"
	"nexus-seed-bot/api"
	"time"

	"nexus-seed-bot/handler"

	seedUtils "nexus-seed-bot/utils"

	nexuslogger "github.com/Nexus-Telegram/nexus/logger"
	"github.com/Nexus-Telegram/nexus/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New().SetTimeout(30 * time.Second)
	client.SetBaseURL("https://alb.seeddao.org/api/v1")
	logger, err := nexuslogger.NewLogger(false)
	if err != nil {
		panic(err)
	}
	queryIDs := utils.ParseQueryIDs()

	logger.Info(fmt.Sprintf("found %d query ids in ", len(queryIDs)))
	for _, queryID := range queryIDs {

		authHeaders := map[string]string{
			"Telegram-Data": queryID,
		}
		headers := utils.MergeHeaders(seedUtils.GetCommonHeaders(), authHeaders)
		client.SetHeaders(headers)

		service := api.NewService(client, logger)
		handler.HandleSeedClaim(service)

		//handler.HandleInitialize(service)
		//go handler.HandleWormCatching(service)
		//go handler.HandleUpgrade(service)
		//go handler.HandleTasks(service)

	}
	select {}
}
