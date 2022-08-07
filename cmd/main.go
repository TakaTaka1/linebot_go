package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"net/url"
	"strings"
	"context"
	
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
    Name string `json:"Name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("lambda finished! response will be returned!")
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	accessToken := os.Getenv("LINE_ACCESS_TOKEN_TEST")
	msg := "テストメッセージ"	
	URL := "https://notify-api.line.me/api/notify"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println("send error")
		log.Fatal(err)
	}

	c := &http.Client{}

	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		fmt.Println("send error")
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, err := c.Do(req)
	if err != nil {
		fmt.Println("send error")
		log.Fatal(err)
	}
	fmt.Println(res)

	log.Printf("lambda started!")
    lambda.Start(HandleRequest)
}
