package types

import "time"

type ProfileResponse struct {
	Data struct {
		Id               string        `json:"id"`
		Name             string        `json:"name"`
		TgId             int64         `json:"tg_id"`
		IsPremium        bool          `json:"is_premium"`
		Upgrades         []interface{} `json:"upgrades"`
		LastClaim        time.Time     `json:"last_claim"`
		ReferrerId       string        `json:"referrer_id"`
		GiveFirstEgg     bool          `json:"give_first_egg"`
		WalletAddressTon interface{}   `json:"wallet_address_ton"`
		Status           struct {
			Active bool `json:"active"`
		} `json:"status"`
		Age                int  `json:"age"`
		TopRate            int  `json:"top_rate"`
		BonusClaimed       bool `json:"bonus_claimed"`
		Achieve            bool `json:"achieve"`
		AchieveFriendBadge bool `json:"achieve_friend_badge"`
		AchieveWeb3Badge   bool `json:"achieve_web3_badge"`
		HalloweenTheme     bool `json:"halloween_theme"`
	} `json:"data"`
}
type Metadata struct {
	URL             string            `json:"url"`
	Name            string            `json:"name"`
	ImageURL        string            `json:"image_url"`
	GroupName       string            `json:"group_name"`
	GroupOrder      int               `json:"group_order"`
	PremiumOnly     bool              `json:"premium_only,omitempty"`
	QueryURL        string            `json:"query_url,omitempty"`
	SubgroupImg     string            `json:"subgroup_img,omitempty"`
	SubgroupName    string            `json:"subgroup_name,omitempty"`
	SubgroupType    string            `json:"subgroup_type,omitempty"`
	StoryMedia      string            `json:"story_media,omitempty"`
	StoryContent    string            `json:"story_content,omitempty"`
	AuthHeaderEnv   string            `json:"auth_header_env,omitempty"`
	AuthHeaderName  string            `json:"auth_header_name,omitempty"`
	AuthHeaderValue string            `json:"auth_header_value,omitempty"`
	QueryMethod     string            `json:"query_method,omitempty"`
	LookupValue     string            `json:"lookup_value,omitempty"`
	QueryParams     map[string]string `json:"query_params,omitempty"`
	Texts           []string          `json:"texts,omitempty"`
	IosURL          string            `json:"ios_url,omitempty"`
	Once            bool              `json:"once,omitempty"`
}

type Task struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	RewardAmount int64    `json:"reward_amount"`
	Sort         int      `json:"sort"`
	Metadata     Metadata `json:"metadata"`
	Repeats      int      `json:"repeats"`
	Tickets      int      `json:"tickets"`
	TaskUser     *string  `json:"task_user"`
}

type TaskResponse struct {
	Data []Task `json:"data"`
}

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
type CatchData struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	Reward    int64     `json:"reward"`
	OnMarket  bool      `json:"on_market"`
	OwnerID   string    `json:"owner_id"`
}

type CatchResponse struct {
	Data CatchData `json:"data"`
}
