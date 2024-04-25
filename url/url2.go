package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bregydoc/gtranslate"
)

// 오늘의 명언을 가져오기 위한 API URL
const quoteURL = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en"

// 오늘의 명언을 나타내는 구조체
type QuoteResponse struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

// 오늘의 명언을 가져오는 함수
func getQuote() (string, string, error) {
	response, err := http.Get(quoteURL)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	var quoteResp QuoteResponse
	err = json.NewDecoder(response.Body).Decode(&quoteResp)
	if err != nil {
		return "", "", err
	}

	return quoteResp.QuoteText, quoteResp.QuoteAuthor, nil
}

// 영어를 한글로 번역하는 함수
func translateToKorean(text string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ko",
		},
	)
	if err != nil {
		return "", err
	}
	return translated, nil
}

func url2() {
	// 오늘의 명언 가져오기
	quote, author, err := getQuote()
	if err != nil {
		fmt.Println("Failed to get today's quote:", err)
		return
	}

	// 영어 명언을 한글로 번역
	translatedQuote, err := translateToKorean(quote)
	if err != nil {
		fmt.Println("Failed to translate quote:", err)
		return
	}

	// 오늘의 명언을 텔레그램 채널로 전송
	message := fmt.Sprintf("오늘의 명언\n\n%s\n\n- %s -", translatedQuote, author)

	fmt.Println("Quote sent to Telegram successfully!", message)
}
