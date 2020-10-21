package main

import (
	"net/url"
	"log"
	"os"
	"net/http"
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", d)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func d(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	consumer_key := os.Getenv("CONSUMER_KEY")
	consumer_secret := os.Getenv("CONSUMER_SECRET")
	accsess_token := os.Getenv("ACCESS_TOKEN")
	accsess_token_secret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(accsess_token, accsess_token_secret)

	v := url.Values{}
	v.Set("count", "200")

	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		if tweet.InReplyToStatusID == 0 {
			api.DeleteTweet(tweet.Id, true)
		}
	}
}