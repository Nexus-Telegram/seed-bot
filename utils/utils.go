package utils

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func GetCommonHeaders() map[string]string {
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

func WaitUntilNextWorm(nextWormTime time.Time, logger *zap.Logger) {
	durationUntilNextWorm := time.Until(nextWormTime.Add(1 * time.Minute))
	if durationUntilNextWorm > 0 {
		logger.Info(fmt.Sprintf("Waiting until %s to catch the worm.", nextWormTime))
		time.Sleep(durationUntilNextWorm)
		logger.Info("It's time to catch the worm!")
	} else {
		logger.Info("The next worm time has already passed.")
	}
}
