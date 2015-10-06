# golack

**娯楽で作ったgolack**

## Description

Simple Incoming Webhooks in Go

## Usage

```go
package main

import "github.com/operando/golack"

func main() {
	slack := golack.Slack{
		Text:      "@operando: test",
		Username:  "sushi",
		IconEmoji: ":sushi:",
		Channel:   "#general",
		LinkNames: true,
	}
	p := golack.Payload{
		slack,
	}
	w := golack.Webhook{
		"webhook_url", // your Incoming WebHooks URL for Slack
	}
	golack.Post(p, w)
}
```

## Install

To install, use `go get`:

```bash
go get github.com/operando/golack
```
