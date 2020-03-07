package auth

import (
	"fmt"

	"github.com/mrjones/oauth"
)

func Authorize(consumerKey, consumerSecret string) (string, string, error) {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return "", "", fmt.Errorf("consumer.GetRequestTokenAndUrl failed: %w", err)
	}

	var pin string
	fmt.Println(url)
	fmt.Scanln(&pin)

	accessToken, err := consumer.AuthorizeToken(requestToken, pin)
	if err != nil {
		return "", "", fmt.Errorf("consumer.AuthorizeToken failed: %w", err)
	}

	return accessToken.Token, accessToken.Secret, nil
}
