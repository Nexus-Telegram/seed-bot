package types

import "time"

type CatchMetadata struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	EndedAt   time.Time `json:"ended_at"`
	NextWorm  time.Time `json:"next_worm"`
	Reward    int64     `json:"reward"`
	IsCaught  bool      `json:"is_caught"`
}

type CatchMetadataResponse struct {
	Data CatchMetadata `json:"data"`
}
type CatchedWorm struct {
	Data struct {
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		Status    string    `json:"status"`
		UpdatedAt time.Time `json:"updated_at"`
		Reward    int       `json:"reward"`
		OnMarket  bool      `json:"on_market"`
		OwnerId   string    `json:"owner_id"`
	} `json:"data"`
}
