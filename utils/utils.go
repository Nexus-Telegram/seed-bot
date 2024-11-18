package utils

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
