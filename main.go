package main

import (
	"log"
	"os"

	"github.com/AM-Myrick/RandomAnimalBot/requests"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	client := createTwitterClient()
	requests.GetDogPic()
	listenForMentions(client)
}

func pickRandomAPI(tweet *twitter.Tweet, client *twitter.Client) {
	requests.GetDogPic()
}

func createTwitterClient() *twitter.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	return client
}

func listenForMentions(client *twitter.Client) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if tweet.InReplyToStatusID != 0 {
			// retweet the status id tweet
		} else {
			// fmt.Println(tweet.User.ScreenName)
			pickRandomAPI(tweet, client)
		}
	}

	params := &twitter.StreamFilterParams{
		Track:         []string{"@RandomAnimalBot"},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(params)

	if err != nil {
		log.Fatal("Failed to open stream.")
	}

	demux.HandleChan(stream.Messages)
}
