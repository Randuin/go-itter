package main

import (
	"fmt"

	a "github.com/ChimeraCoder/anaconda"
	c "github.com/ttacon/chalk"
)

func PrintTweet(t a.Tweet) {
	fmt.Println(c.Red, t.User.ScreenName, c.Reset, t.Text)
}

func main() {
	a.SetConsumerKey("")
	a.SetConsumerSecret("")
	api := a.NewTwitterApi("", "")
	tweet_stream := api.UserStream(nil)
	for {
		event := <-tweet_stream.C
		switch event.(type) {
		case a.Tweet:
			PrintTweet(event.(a.Tweet))
		}
	}
}
