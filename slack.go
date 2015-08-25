package golack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ParameterPayload = "payload"
)

type Slack struct {
	Text       string `json:"text"`       //投稿内容
	Username   string `json:"username"`   //投稿者名 or Bot名（存在しなくてOK）
	Icon_emoji string `json:"icon_emoji"` //アイコン絵文字
	Icon_url   string `json:"icon_url"`   //アイコンURL（icon_emojiが存在する場合は、適応されない）
	Channel    string `json:"channel"`    //#部屋名
}

type Payload struct {
	IncomingUrl string
	Slack       Slack
}

func post(payload Payload) {
	params, _ := json.Marshal(payload.Slack)
	resp, _ := http.PostForm(
		payload.IncomingUrl,
		url.Values{ParameterPayload: {string(params)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	println(string(body))
}
