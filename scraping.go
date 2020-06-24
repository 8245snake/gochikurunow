package gochikurunow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//getToken tokenを取得する
func getToken(client *http.Client) string {
	resp, err := client.Get(LoginArrd)
	if err != nil {
		fmt.Println("[Error]GetSessionID client.Do failed", err)
		return ""
	}
	defer resp.Body.Close()

	doc, e := goquery.NewDocumentFromResponse(resp)
	if e != nil {
		fmt.Println("[Error]GetSessionID NewDocumentFromResponse failed", e)
		return ""
	}

	token, _ := doc.Find("input[name=_token]").First().Attr("value")
	return token
}

//createPostRequestJSON リクエスト作成
func createPostRequestJSON(mail, password, token string) *http.Request {
	//リクエストBody作成
	type Forms struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"_token,omitempty"`
	}

	frm := Forms{Email: mail, Password: password, Token: token}
	marshalized, err := json.Marshal(frm)
	if err != nil {
		fmt.Println("[Error]CreatePostRequestJSON create NewRequest failed", err)
		return nil
	}
	req, err := http.NewRequest(
		"POST",
		LoginArrd,
		bytes.NewBuffer(marshalized),
	)
	if err != nil {
		fmt.Println("[Error]CreatePostRequestJSON create NewRequest failed", err)
		return nil
	}
	// リクエストHead作成
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://gochikurunow.com/login")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36")

	return req
}

//postRequest ポスト
func postRequest(req *http.Request, client *http.Client) (*goquery.Document, error) {
	if req == nil {
		return nil, fmt.Errorf("error")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Error]GetSessionID client.Do failed", err)
		return nil, err
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromResponse(resp)
}

//extractDateFromTitleString 日付を抽出
func extractDateFromTitleString(title string) string {
	delimiter := "||"
	base := strings.Replace(title, "<small>", delimiter, -1)
	base = strings.Replace(base, "</small>", delimiter, -1)
	baseArr := strings.Split(base, delimiter)
	return strings.TrimSpace(baseArr[0]) +
		strings.TrimSpace(baseArr[1]) +
		strings.TrimSpace(baseArr[2]) +
		strings.TrimSpace(baseArr[3]) +
		strings.TrimSpace(baseArr[4])

}

//extractPrice 金額を抽出
func extractPrice(raw string) string {
	price := strings.Replace(raw, "円（税込）", "", -1)
	return strings.TrimSpace(price)

}
