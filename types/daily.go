package types

import "time"

type Status string

const (
	Received Status = "received"
	Created  Status = "created"
)

type StreakRewards struct {
	Data []struct {
		Id        string    `json:"id"`
		UserId    string    `json:"user_id"`
		Status    Status    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}
type StreakRewardsPayload struct {
	Ids []string `json:"streak_reward_ids"`
}
type LoginStreak struct {
	Data struct {
		No        int       `json:"no"`
		CreatedAt time.Time `json:"created_at"`
		NoMax     int       `json:"no_max"`
	} `json:"data"`
}
