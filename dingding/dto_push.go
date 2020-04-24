package dingding

type PushText struct {
	Content string `json:"content"`
}

type PushAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type PushRequest struct {
	MsgType string   `json:"msgtype"`
	Text    PushText `json:"text"`
	At      PushAt   `json:"at"`
}
