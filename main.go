package main

// Source code based on the example:
// https://github.com/michimani/gotwi/blob/main/_examples/2_post_delete_tweet/main.go

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func main() {

	// Check expected secrets are set in the environment variables
	accessToken := os.Getenv("TW_ACCESS_TOKEN")
	accessSecret := os.Getenv("TW_ACCESS_SECRET")

	if accessToken == "" || accessSecret == "" {
		fmt.Fprintln(os.Stderr, "Please set the TW_ACCESS_TOKEN and TW_ACCESS_SECRET environment variables.")
		os.Exit(1)
	}

	// Check the comment to tweet has been supplied
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Please pass the comment to tweet as the only argument.")
		os.Exit(1)
	}

	client, err := newOAuth1Client(accessToken, accessSecret)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tweetId, err := tweet(client, os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println("tweet id", tweetId)
}

func newOAuth1Client(accessToken, accessSecret string) (*gotwi.Client, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           accessToken,
		OAuthTokenSecret:     accessSecret,
	}

	return gotwi.NewClient(in)
}

func tweet(c *gotwi.Client, text string) (string, error) {
	p := &types.CreateInput{
		Text: gotwi.String(text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}
