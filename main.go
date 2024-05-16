package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/go-github/v60/github"
	"golang.org/x/text/message"
	"net/http"
)

const (
	awtrixURL = "http://192.168.88.74/api/notify"
)

type Payload struct {
	Duration int    `json:"duration"`
	Rainbow  bool   `json:"rainbow"`
	Text     string `json:"text"`
}

func main() {
	// read repo info
	ghClient := github.NewClient(nil)
	repo, _, err := ghClient.Repositories.Get(context.Background(), "kubescape", "kubescape")
	if err != nil {
		panic(err)
	}
	// send to awtrix
	count := repo.GetStargazersCount()
	var rainbow bool
	if count%100 == 0 {
		rainbow = true
	}
	p := message.NewPrinter(message.MatchLanguage("en"))
	payload := Payload{
		Duration: 10,
		Rainbow:  rainbow,
		Text:     p.Sprintf("%d", repo.GetStargazersCount()),
	}
	jsonData, err := json.Marshal(payload)
	request, err := http.NewRequest("POST", awtrixURL, bytes.NewBuffer(jsonData))
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		panic(err)
	}
}
