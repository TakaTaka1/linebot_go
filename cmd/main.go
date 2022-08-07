package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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

func LineNotify () {
	f := "./.env"
	if _, err := os.Stat(f); err == nil {
		err_read := godotenv.Load(f)
		if err_read != nil {
			log.Fatalf("error: %v", err_read)
		}
		fmt.Println(".env is existed")
	} else {
		fmt.Println(".env is not existed")
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
}

func main() {
	// LineNotify()
	// f := "./.env"
	// if _, err := os.Stat(f); err == nil {
	// 	err_read := godotenv.Load(f)
	// 	if err_read != nil {
	// 		log.Fatalf("error: %v", err_read)
	// 	}
	// 	fmt.Println(".env is existed")
	// } else {
	// 	fmt.Println(".env is not existed")
	// }

	// accessToken := os.Getenv("LINE_ACCESS_TOKEN_TEST")
	// msg := "テストメッセージ"
	// URL := "https://notify-api.line.me/api/notify"

	// u, err := url.ParseRequestURI(URL)
	// if err != nil {
	// 	fmt.Println("send error")
	// 	log.Fatal(err)
	// }

	// c := &http.Client{}

	// form := url.Values{}
	// form.Add("message", msg)

	// body := strings.NewReader(form.Encode())

	// req, err := http.NewRequest("POST", u.String(), body)
	// if err != nil {
	// 	fmt.Println("send error")
	// 	log.Fatal(err)
	// }

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Authorization", "Bearer "+accessToken)

	// res, err := c.Do(req)
	// if err != nil {
	// 	fmt.Println("send error")
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)
	
	log.Printf("lambda started!")
	// lambda.Start(HandleRequest)
	lambda.Start(LineNotify)
}
