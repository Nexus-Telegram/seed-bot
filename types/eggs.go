package types

import "time"

type EggData struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Items    []struct {
		Id      string `json:"id"`
		OwnerId string `json:"owner_id"`
		Status  string `json:"status"`
		Type    string `json:"type"`
	} `json:"items"`
}

type GetMyEggs struct {
	Data EggData `json:"data"`
}
type FirstEggData struct {
	Id      string `json:"id"`
	OwnerId string `json:"owner_id"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}
type FirstEgg struct {
	FirstEggData `json:"data"`
}
type HatchEgg struct {
	Data struct {
		Id               string      `json:"id"`
		Type             string      `json:"type"`
		Status           string      `json:"status"`
		EnergyLevel      int         `json:"energy_level"`
		EnergyMax        int         `json:"energy_max"`
		HappinessLevel   int         `json:"happiness_level"`
		TaskLevel        int         `json:"task_level"`
		IsLeader         bool        `json:"is_leader"`
		HuntStartAt      time.Time   `json:"hunt_start_at"`
		HuntEndAt        time.Time   `json:"hunt_end_at"`
		ReceivedRewardAt time.Time   `json:"received_reward_at"`
		OnMarket         bool        `json:"on_market"`
		OwnerId          string      `json:"owner_id"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		MarketId         interface{} `json:"market_id"`
		Price            interface{} `json:"price"`
	} `json:"data"`
}
