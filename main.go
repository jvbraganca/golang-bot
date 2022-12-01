package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env File not found")
	}

	config := oauth1.NewConfig(os.Getenv("API_KEY"),
		os.Getenv("API_KEY_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#golang", // the word of phrase we want to query, note that it is case insensitive
		Count: 5,         // the amount of tweets to be returned
	})

	// check for errors and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// iterates over the results and retweet every tweet
	for _, tweet := range tweets.Statuses {
		_, _, err := client.Statuses.Retweet(tweet.ID, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
