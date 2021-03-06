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
	Text      string `json:"text" toml:"text"`             //投稿内容
	Username  string `json:"username" toml:"username"`     //投稿者名 or Bot名（存在しなくてOK）
	IconEmoji string `json:"icon_emoji" toml:"icon_emoji"` //アイコン絵文字
	IconURL   string `json:"icon_url" toml:"icon_url"`     //アイコンURL（icon_emojiが存在する場合は、適応されない）
	Channel   string `json:"channel" toml:"channel"`       //#部屋名
	LinkNames bool   `json:"link_names" toml:"link_names"`
}

type Webhook struct {
	URL string `toml:"url"`
}

type Payload struct {
	Slack Slack
}

func Post(payload Payload, webhook Webhook) {
	params, _ := json.Marshal(payload.Slack)
	resp, _ := http.PostForm(
		webhook.URL,
		url.Values{ParameterPayload: {string(params)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	println(string(body))
}
