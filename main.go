package main

import (
	"fmt"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/headzoo/surf/errors"
)

const (
	zerodhaLoginURL   = "https://zerodha.com/login"
	googleAuthPageURL = "https://accounts.google.com/signin"
)

func main() {
	// Create a new browser instance
	b := surf.NewBrowser()

	// Open Zerodha login page
	err := b.Open(zerodhaLoginURL)
	if err != nil {
		fmt.Println("Failed to open Zerodha login page:", err)
		return
	}

	// Find the login form
	fm, err := b.Form("form[action='/login/post']")
	if err != nil {
		fmt.Println("Failed to find login form:", err)
		return
	}

	// Enter Zerodha username and password
	fm.Input("user_id", "rahulchako7@gmail.com")
	fm.Input("password", "Rahulchacko@1996")

	// Submit login form
	err = fm.Submit()
	if err != nil {
		fmt.Println("Failed to submit login form:", err)
		return
	}

	// Check if redirected to Google authentication
	if b.Url().String() == googleAuthPageURL {
		fmt.Println("Performing Google authentication...")
		// Handle Google authentication here
	}

	// Extract or generate authorization token
	authToken, err := extractAuthToken(b)
	if err != nil {
		fmt.Println("Failed to extract authorization token:", err)
		return
	}

	fmt.Println("Authorization Token:", authToken)
}

func extractAuthToken(b *browser.Browser) (string, error) {
	// Example of extracting authorization token from response headers
	for _, cookie := range b.SiteCookies() {
		if cookie.Name == "authorization_token" {
			return cookie.Value, nil
		}
	}
	return "", errors.New("authorization token not found")
}
