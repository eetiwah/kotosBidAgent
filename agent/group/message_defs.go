package group

type GroupMessage struct {
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

type GroupInvite struct {
	GroupID   string `json:"GroupID"`
	GroupName string `json:"GroupName"`
	//	SignedGroupID string `json:"SignedGroupID"` // Nullable string
	Timestamp  int    `json:"Timestamp"`
	SharedKey  string `json:"SharedKey"`
	ServerHost string `json:"ServerHost"`
}
