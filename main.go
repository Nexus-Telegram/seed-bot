package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	nexusUtils "github.com/nexus-telegram/nexus-core/utils"
	"github.com/nexus-telegram/seed-bot/api"
	"github.com/nexus-telegram/seed-bot/utils"

	"github.com/nexus-telegram/seed-bot/handler"

	"github.com/go-resty/resty/v2"
	nexuslogger "github.com/nexus-telegram/nexus-core/logger"
)

func authenticate(s *api.Service) {
	res, err := s.Client.R().Post("/profile")
	if err != nil {
		s.Logger.Error("Request failed: " + err.Error())
		return
	}

	// Check for HTTP status 400 (Bad Request)
	if res.StatusCode() == 400 {
		var responseData map[string]string

		// Try to unmarshal the response body into a map
		if err := json.Unmarshal(res.Body(), &responseData); err != nil {
			s.Logger.Error("Failed to parse error response: " + err.Error())
			return
		}

		// Check if the error code indicates an existing user
		if responseData["code"] == "invalid-request" && responseData["message"] == "user with telegram id already exist" {
			s.Logger.Debug("User is already connected with this Telegram ID")
		} else {
			s.Logger.Info("Received error: " + responseData["message"])
		}
		return
	}
	s.Logger.Debug("User connected successfully")
}

func processAccount(queryID string) {
	client := resty.New().SetTimeout(30 * time.Second)
	client.SetBaseURL("https://alb.seeddao.org/api/v1")
	proxyHost := "rp.proxyscrape.com:6060"
	username := "xjfxgi690pe0dkx"
	password := "5c3oqgl5yu263gk"
	proxyURL := fmt.Sprintf("http://%s:%s@%s", username, password, proxyHost)
	parsedProxyURL, err := url.Parse(proxyURL)
	if err != nil {
		panic("Invalid proxy URL")
	}
	client.SetTransport(&http.Transport{
		Proxy: http.ProxyURL(parsedProxyURL),
	})
	resp, err := client.R().
		Get("http://httpbin.org/ip")

	if err != nil {
		fmt.Println("Proxy failed", err.Error())
	} else {
		fmt.Println("Proxy working", resp.String())
	}
	logger := nexuslogger.New()
	if logger == nil {
		panic(err)
	}
	authHeaders := map[string]string{
		"Telegram-Data": queryID,
	}
	headers := nexusUtils.MergeHeaders(utils.GetCommonHeaders(), authHeaders)
	client.SetHeaders(headers)

	service := api.NewService(client, logger)
	authenticate(service)

	go handler.HandleInitializeBird(service)
	go handler.HandleDaily(service)
	go handler.HandleSeedClaim(service)
	go handler.HandleWormCatching(service)
	// go handler.HandleUpgrade(service)
	go handler.HandleTasks(service)
	go handler.HandleBird(service)
}

func main() {

	queryIDs := nexusUtils.ParseQueryIDs()

	for _, queryID := range queryIDs {
		go processAccount(queryID)
		time.Sleep(1 * time.Second)
	}
	select {}
}
