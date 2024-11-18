package types

import "time"

type Upgrade struct {
	UpgradeType  string    `json:"upgrade_type"`
	UpgradeLevel int       `json:"upgrade_level"`
	Timestamp    time.Time `json:"timestamp"`
}
type ProfileResponse struct {
	Data struct {
		Id               string      `json:"id"`
		Name             string      `json:"name"`
		TgId             int64       `json:"tg_id"`
		IsPremium        bool        `json:"is_premium"`
		Upgrades         []Upgrade   `json:"upgrades"`
		LastClaim        time.Time   `json:"last_claim"`
		ReferrerId       string      `json:"referrer_id"`
		GiveFirstEgg     bool        `json:"give_first_egg"`
		WalletAddressTon interface{} `json:"wallet_address_ton"`
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

type Progress struct {
	Data []struct {
		Id           string `json:"id"`
		Type         string `json:"type"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		RewardAmount int64  `json:"reward_amount"`
		Sort         int    `json:"sort"`
		Metadata     struct {
			Url         string      `json:"url,omitempty"`
			Name        string      `json:"name,omitempty"`
			ImageUrl    string      `json:"image_url"`
			GroupName   string      `json:"group_name,omitempty"`
			GroupOrder  interface{} `json:"group_order,omitempty"`
			PremiumOnly bool        `json:"premium_only,omitempty"`
			IsTma       bool        `json:"is_tma,omitempty"`
			QueryUrl    string      `json:"query_url,omitempty"`
			LookupValue string      `json:"lookup_value,omitempty"`
			QueryMethod string      `json:"query_method,omitempty"`
			QueryParams struct {
				MinGem  string `json:"min_gem"`
				UserId  string `json:"user_id"`
				RefCode string `json:"ref_code"`
			} `json:"query_params,omitempty"`
			SubgroupImg     string   `json:"subgroup_img,omitempty"`
			SubgroupName    string   `json:"subgroup_name,omitempty"`
			SubgroupType    string   `json:"subgroup_type,omitempty"`
			AuthHeaderEnv   string   `json:"auth_header_env,omitempty"`
			AuthHeaderName  string   `json:"auth_header_name,omitempty"`
			AuthHeaderValue string   `json:"auth_header_value,omitempty"`
			StoryMedia      string   `json:"story_media,omitempty"`
			StoryContent    string   `json:"story_content,omitempty"`
			Texts           []string `json:"texts,omitempty"`
			IosUrl          string   `json:"ios_url,omitempty"`
			Once            bool     `json:"once,omitempty"`
			Excluded        bool     `json:"excluded,omitempty"`
			AnswerLength    int      `json:"answer_length,omitempty"`
			Important       string   `json:"important,omitempty"`
		} `json:"metadata"`
		Repeats  int `json:"repeats"`
		Tickets  int `json:"tickets"`
		TaskUser *struct {
			Id           string `json:"id"`
			Completed    bool   `json:"completed"`
			RewardAmount int    `json:"reward_amount"`
			Repeats      int    `json:"repeats"`
			Tickets      int    `json:"tickets"`
		} `json:"task_user"`
	} `json:"data"`
}
type Settings struct {
	Data struct {
		BoardingEvent struct {
			From  time.Time `json:"from"`
			Scale int       `json:"scale"`
			To    time.Time `json:"to"`
		} `json:"boarding_event"`
		EnergyMaxBird struct {
			Hawk    int64 `json:"hawk"`
			Owl     int64 `json:"owl"`
			Parrot  int64 `json:"parrot"`
			Penguin int64 `json:"penguin"`
			Phoenix int64 `json:"phoenix"`
			Sparrow int   `json:"sparrow"`
		} `json:"energy-max-bird"`
		Fee       int `json:"fee"`
		HappyDays struct {
		} `json:"happy-days"`
		HappyDaysRewards []struct {
			Id          string `json:"id"`
			Type        string `json:"type"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Amount      int64  `json:"amount"`
			Unit        string `json:"unit"`
			ExpiredIn   int    `json:"expired_in"`
			Weight      int    `json:"weight"`
		} `json:"happy-days-rewards"`
		HolyWater            []int     `json:"holy-water"`
		HolyWaterCosts       []int     `json:"holy-water-costs"`
		HuntingTimeline      time.Time `json:"hunting-timeline"`
		LoginBonuses         []int     `json:"login-bonuses"`
		MiningSpeed          []int     `json:"mining-speed"`
		MiningSpeedCosts     []int64   `json:"mining-speed-costs"`
		StorageSize          []int     `json:"storage-size"`
		StorageSizeCosts     []int64   `json:"storage-size-costs"`
		TransferWormToEnergy struct {
			Common    int   `json:"common"`
			Epic      int64 `json:"epic"`
			Legendary int64 `json:"legendary"`
			Rare      int64 `json:"rare"`
			Seed      int64 `json:"seed"`
			Uncommon  int64 `json:"uncommon"`
		} `json:"transfer-worm-to-energy"`
	} `json:"data"`
}

type TaskResponse struct {
	Data []Task `json:"data"`
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
type Balance struct {
	Balance int `json:"data"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
