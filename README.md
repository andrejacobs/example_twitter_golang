# example_twitter_golang

Send a tweet to Twitter using Go.

This example code is a precursor for me while I am working on my hobby project [ajtweet](https://github.com/andrejacobs/ajtweet-cli/) that I will be using to schedule tweets.

See my [blog post](TODO) for more details on how to explore the Twitter APIs and how I wrote this code.

**NOTE:** This uses the [gotwi library](github.com/michimani/gotwi) for doing the heavy work.

## Configuration

Install the required go packages:

    go mod tidy

Configure your Twitter API Key, Secret and OAuth 1 User token and secret.

    cp env_example .env
    # Edit the .env file.

## Build and Send a tweet

Build the executable:

    go build -o tweet .

Source the environment variables:

    source .env

Send a tweet:

    ./tweet "Test: Hello world from golang"
    tweet id 1527644512285343745
