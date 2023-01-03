## HOW TO MAKE A TWITTER BOT WITH GO LANG

As a software engineer, I am always looking for ways to automate things using APIs.
And having a bot doing what you wish is really cool.
For instance, if you want to share some knowledge with your Twitter community you can write a Bot that searches for the
information and publish it to you without any sweat. So today we are going to learn how to write, query,
and retweet using the Twitter API and a GoLang client.

### Creating your Bot

The first step you need to take for creating a Twitter Bot is to sign up for a developer account at
https://developer.twitter.com/. Log in with your Twitter account and press the Sign up button at the top right corner 
of your screen. Make sure you are logged in to the account you wish to automate.

![](./readme-assets/twitter-dev-menu.png)







There, you will click on the `Create Project` button. Here you will describe your project. First, enter the project
name and then the use case. For this article, it was prompted by the following:

1. Project name: GoLand Bot
1. Use case: Making a bot
1. Project description: A bot to tweet, retweet, and reply to others.

The App environment selected was development,

![](./readme-assets/choose-env-menu.png)

Then we select the app name where we will be able to get our access token and key.

![](./readme-assets/name-app-tab.png)

















On the next page, your API Key and API Key Secret will be displayed, copy them to a safe place. And then click on the `Go To Dashboard` button. We need to configure our app authentication for this click on the `gear icon` and go to the `App Settings` page.

![](./readme-assets/app-settings-btn.png)









Under `User authentication settings` click on the `Set up` button

![](./readme-assets/user-auth-settings.png)











Under `App Permissions` we `MUST` give `Read and write` permissions. The `Type of App` is `Web App, Automated App or Bot`.

![](./readme-assets/user-auth-settings-permissions.png)

![](./readme-assets/type-app-menu.png)

Under `App info` feel free to insert whatever you want, under `callback URI` and `Website URL` I inserted the bot’s GitHub repository. Then, click on the `Save` button. You will then be prompted your `Client ID` and `Client Secret`, copy them if you want to use `OAuth 2`, for this article we will be using `OAuth 1`.

![](./readme-assets/app-info-form.png)

When redirected to the project dashboard, go to the `Keys and tokens` tab and copy your `access token` and `secret`. Make sure it was generated with `Read and Write` permissions.

![](./readme-assets/application-summary.png)

Again, copy them to a safe place. Now we are ready to begin coding.

## Creating the GoLang project

Open your terminal and create a Go project. First, create a folder to our project, mine will be named `golang-bot`, and change it to its directory.

```bash
 mkdir golang-bot
```
```bash
 cd golang-bot
```

Now, we can create our GoLang project. Move into your favorite IDE, mine is GoLand. Create a `main.go` file so we can write the bot’s code, and initialize a module with
```bash
 go mod init theNameOrPathToYourModule
```

## Installing the necessary libraries

In this project we have three necessary libs. Install them with

```bash
go get github.com/joho/godotenv
go get github.com/dghubble/go-twitter
go get github.com/dghubble/oauth1
```

## Configuring the project

Then, create your `.env` file and populate them with the `API KEY, API KEY SECRET, ACCESS TOKEN, ACCESS TOKEN SECRET`, in the following format:

```env
API_KEY=
API_KEY_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```
## Loading the environment variables

We proceed to load our environment variables with `godotenv` package, for that we must load them using `godotenv.Load(“.env”)`

```go
err := godotenv.Load(".env")
// if there is any problem loading the variables the program will exit
if err != nil {
   log.Fatalf(".env File not found")
}
```

Now we want to configure our access to the Twitter through `oauth1` package. For that, do the following:

```go
// setting up the oauth1 library with our api key and secret
config := oauth1.NewConfig(os.Getenv("API_KEY"), os.Getenv("API_KEY_SECRET"))
// setting up the token
token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
```

## Creating the client

Now it’s necessary to create a httpClient with `oauth1` and then a `Twitter Client.`

```go
// creating a HTTP Client with the token
httpClient := config.Client(oauth1.NoContext, token)
// creates a twitter client from the installed package with a httpClient
client := twitter.NewClient(httpClient)
```

## Posting a tweet from the bot

Finally we are able to have some fun. To create a tweet you can use the `client` we declared with the `twitter` package. And for tweeting you may code:

```go
tweet, res, err := client.Statuses.Update("Hello, World! This is the first message from the twitter bot.", nil)

// check for errors and if there is any it will print it and exit the program
if err != nil {
  log.Fatal(err)
}
```

The `res` variable contains the `http.Response` from our request to the `Twitter Api`, in this case it is irrelevant, so I going to replace it with `_.`  The `tweet` variable contains all the data of the posted tweet. To see what was posted, you may log it with `tweet.Text`

```go
log.Println(tweet.Text)
```

Now, run the program on the terminal with

```bash
go run main.go
```
And here is the output

```bash
2022/12/01 20:18:40 Hello, World! This is the first message from the twitter bot.
```

You can see the tweet by going to your Twitter profile, here is mine:

![](./readme-assets/tweet-result.png)

## Querying tweets

GREAT! Now you know how to post tweets from your won client. The next step is to learn how to search tweets containing certain keywords. For this, we will still be using the same Twitter client but a different method. If we want to search for tweets containing the hashtag `#golang` all we have to do is

```go
tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
  Query: "#golang", // the word of phrase we want to query, note that it is case insensitive
  Count: 5,         // the amount of tweets to be returned
})

// check for errors and exit the program
if err != nil {
  log.Fatal(err)
}

// iterates over the results and print them
for _, tweet := range tweets.Statuses {
  log.Print(tweet.Text)
}
```

## Retweeting

And run the program again to see the results. The last thing for us to learn is how to retweet. It is pretty simple as well. After querying for a list of tweets all you have to do is:


```go
// iterates over the results and retweet every tweet
for _, tweet := range tweets.Statuses {
  _, _, err := client.Statuses.Retweet(tweet.ID, nil)
  if err != nil {
     log.Fatal(err)
  }
}
```

## Whole content of main.go

```go
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
```

## Conclusion

Here you found the basics to start using the Twitter API. You may also create a cron job to run this bot every 15 minutes, or host it on Azure Functions, and it will keep running and executing your tasks. If you are eager to discover what else you are able to do it the `Go Twitter` package visit <https://pkg.go.dev/github.com/dghubble/go-twitter/twitter> to learn more. You can check out the whole bot code at: <https://github.com/jvbraganca/golang-bot>.
