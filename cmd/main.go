package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func GetWeekNumber(year, month, day int) int {
	CurrentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekNum := float64(CurrentDate.Day()) / 7

	return int(math.Ceil(weekNum))
}

type MyEvent struct {
	Name string `json:"Name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("lambda finished! response will be returned!")

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	f := "./.env"
	err_read := godotenv.Load(f)
	if err_read != nil {
		log.Fatalf("error: %v", err_read)
	}
	host, _ := os.Hostname()
	matchedLocal, _ := regexp.Match(`local`, []byte(host))

	var accessToken string
	accessToken = os.Getenv("LINE_ACCESS_TOKEN_TEST")
	if !matchedLocal {
		accessToken = os.Getenv("LINE_ACCESS_TOKEN")
	}
	// ゴミの日のポストメッセージ
	// ゴミ出し日の前日9pmと当日の7:30amにポストする
	// TODO 曜日判定メソッドのリファクタリング
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
		Kanen     = "🔥可燃ゴミ🔥"
		Funen     = "👀不燃ゴミ👀"
		Shigen    = "♻️資源ゴミ️️️️️️♻️"
	)

	today := time.Now()          // 現在日時の取得
	DayOfWeek := today.Weekday() // 曜日の取得
	msg := ""

	if DayOfWeek == Tuesday {
		// 可燃
		msg += "\n明日は" + Kanen + "収集日です。"
	}
	wN := GetWeekNumber(today.Year(), int(today.Month()), today.Day())

	if DayOfWeek == Wednesday {
		msg += "\n今日は" + Kanen + "収集日です。"
		if wN == 1 || wN == 3 {
			msg += "\n明日は" + Funen + "収集日です。"
		}
	}

	if DayOfWeek == Thursday && (wN == 1 || wN == 3) {
		msg += "\n今日は" + Funen + "収集日です。"
	}

	if DayOfWeek == Thursday {
		// 資源
		msg += "\n明日は" + Shigen + "収集日です。"
	}

	if DayOfWeek == Friday {
		// 資源・可燃ゴミ
		msg += "\n今日は" + Shigen + "収集日です。"
		msg += "\n明日は" + Kanen + "収集日です。"
	}

	if DayOfWeek == Saturday {
		// 資源・可燃ゴミ
		msg += "\n今日は" + Kanen + "収集日です。"
	}

	URL := "https://notify-api.line.me/api/notify"
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println("send error")
		log.Fatal(err)
	}

	c := &http.Client{}

	form := url.Values{}
	if msg == "" {
		os.Exit(0)
	}
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
