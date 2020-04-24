package dingding

import rest "github.com/xiaoguaihou/go-rest"

func Post2Dingding(url string, content string) {

	request := PushRequest{
		MsgType: "text",
		Text: PushText{
			Content: content,
		},
	}

	rest.Post(url, &request, nil)
}
