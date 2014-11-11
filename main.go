package main

import (
	"github.com/go-martini/martini"
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()

	oauthOpts := &gooauth2.Options{
		ClientID:     "cf-go-client-example",
		ClientSecret: "c1oudc0w",
		RedirectURL:  "https://cf-go-client-example.10.244.0.34.xip.io/oauth2callback",
		Scopes:       []string{""},
	}

	cf := oauth2.NewOAuth2Provider(oauthOpts, "https://login.10.244.0.34.xip.io/oauth/authorize",
		"https://uaa.10.244.0.34.xip.io/oauth/token")

	m.Handlers(
		sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))),
		cf,
		oauth2.LoginRequired,
		martini.Logger(),
		martini.Static("public"),
	)

	m.Get("/", func(tokens oauth2.Tokens) string {
		if tokens.IsExpired() {
			return "not logged in, or the access token is expired"
		}
		return "logged in"
	})

	m.Run()
}
