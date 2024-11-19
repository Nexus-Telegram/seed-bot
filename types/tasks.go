package types

type TaskAnswerPayload struct {
	Answer string `json:"answer"`
}

var TaskSecrets = map[string]string{
	"7c7bffa3-cd52-4443-8647-4d2201cb7c7a": "Wallet", // Example mapping
	"52e80660-083a-40df-841c-2a44221dc4b6": "BRESEED",
	"99f3c62f-d396-4f81-8e64-32a972bba2ab": "Bullrun",
	"c39fb00a-2c2e-451b-87ca-23d65e00eacf": "Tokens",
	"ace94ccd-be20-43e3-a105-60dd80717ecb": "Ton",
	"450a4a67-f3bb-4f78-975c-30bba2cb57c3": "BTCTOTHEMOON",
	"ae196f2c-7a8c-47e9-8c86-95f8c27f07b9": "TRANSaction",
	"c29e59ac-55dd-414c-a065-03802e50d3af": "OKXEED",
	"2931da7e-5ac7-4c38-8e44-f52960f4e823": "BIRDIE",
	"a497c5a2-700d-4ea7-ae6d-2b31e84585eb": "GETGEMS",
	"4663208c-551f-4cfe-b95a-63e5deece870": "Airdrop",
}

type TaskType string

const (
	JoinCommunity       TaskType = "Join community"
	PlayApp             TaskType = "Play app"
	TelegramBoost       TaskType = "telegram-boost"
	TGStory             TaskType = "TG story"
	FollowUs            TaskType = "Follow us"
	TelegramNameInclude TaskType = "telegram-name-include"
	AddList             TaskType = "Add_list"
	Refer               TaskType = "refer"
	Academy             TaskType = "academy"
	OKXCommunity        TaskType = "OKX_community"
	OKX                 TaskType = "OKX"
	MintBirdNFT         TaskType = "mint-bird-nft"
	TONWalletConnect    TaskType = "ton-wallet-connect"
)

type QueryParams struct {
	MinGem  string `json:"min_gem"`
	UserId  string `json:"user_id"`
	RefCode string `json:"ref_code"`
}

type TaskMetadata struct {
	Url             string      `json:"url,omitempty"`
	Name            string      `json:"name,omitempty"`
	ImageUrl        string      `json:"image_url"`
	GroupName       string      `json:"group_name,omitempty"`
	GroupOrder      interface{} `json:"group_order,omitempty"`
	PremiumOnly     bool        `json:"premium_only,omitempty"`
	IsTma           bool        `json:"is_tma,omitempty"`
	QueryUrl        string      `json:"query_url,omitempty"`
	LookupValue     string      `json:"lookup_value,omitempty"`
	QueryMethod     string      `json:"query_method,omitempty"`
	QueryParams     QueryParams `json:"query_params,omitempty"`
	SubgroupImg     string      `json:"subgroup_img,omitempty"`
	SubgroupName    string      `json:"subgroup_name,omitempty"`
	SubgroupType    string      `json:"subgroup_type,omitempty"`
	AuthHeaderEnv   string      `json:"auth_header_env,omitempty"`
	AuthHeaderName  string      `json:"auth_header_name,omitempty"`
	AuthHeaderValue string      `json:"auth_header_value,omitempty"`
	StoryMedia      string      `json:"story_media,omitempty"`
	StoryContent    string      `json:"story_content,omitempty"`
	Texts           []string    `json:"texts,omitempty"`
	IosUrl          string      `json:"ios_url,omitempty"`
	Once            bool        `json:"once,omitempty"`
	Excluded        bool        `json:"excluded,omitempty"`
	AnswerLength    int         `json:"answer_length,omitempty"`
	Important       string      `json:"important,omitempty"`
}

type TaskUser struct {
	Id           string `json:"id"`
	Completed    bool   `json:"completed"`
	RewardAmount int    `json:"reward_amount"`
	Repeats      int    `json:"repeats"`
	Tickets      int    `json:"tickets"`
}

type UniqueTask struct {
	Id           string       `json:"id"`
	Type         TaskType     `json:"type"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	RewardAmount int64        `json:"reward_amount"`
	Sort         int          `json:"sort"`
	Metadata     TaskMetadata `json:"metadata"`
	Repeats      int          `json:"repeats"`
	Tickets      int          `json:"tickets"`
	TaskUser     *TaskUser    `json:"task_user"`
}

type Progress struct {
	Data []UniqueTask `json:"data"`
}
