package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SendMessage(message string) {
	config := ReadConfig()

	text := message
	botToken := config.BotToken
	chatId := config.ChatId

	endpoint := "https://api.telegram.org/bot" + botToken + "/sendMessage?" + "chat_id=" + chatId + "&text=" + url.QueryEscape(text)

	_, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	BotToken string `json:"bot_token"`
	ChatId   string `json:"chat_id"`
}

func ReadConfig() Config {

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(file, &config)

	return config
}
