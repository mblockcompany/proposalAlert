package request

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func SendSlackMsg(msg string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("not loaded cuz err",err)
	}
	token := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	api := slack.New(token)
	_, _, err = api.PostMessage(
		channelID,
		slack.MsgOptionText(msg, false),
	)
	if err != nil {
		log.Fatal("err",err)
	}
}