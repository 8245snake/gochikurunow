package gochikurunow

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	//LoginArrd アドレス
	LoginArrd = `https://gochikurunow.com/login`
)

//GochiClient クライアント
type GochiClient struct {
	client  *http.Client
	request *http.Request
}

//NewGochiClient 初期化
func NewGochiClient(mail string, password string) (*GochiClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	client := &http.Client{Jar: jar}
	token := getToken(client)
	if token == "" {
		return nil, fmt.Errorf("ページの取得に失敗しました")
	}
	req := createPostRequestJSON(mail, password, token)
	api := &GochiClient{client: client, request: req}
	return api, nil
}

//GetMenu 本日のメニューを取得
func (api *GochiClient) GetMenu() (MenueInfo, error) {
	var menue MenueInfo
	doc, err := postRequest(api.request, api.client)
	if err != nil {
		fmt.Printf("%v\n", err)
		return menue, err
	}

	head, err := doc.Find("div.productHead").Find("h2").Html()
	if err != nil {
		fmt.Printf("%v\n", err)
		return menue, err
	}
	if head == "" {
		return menue, fmt.Errorf("ページの取得に失敗しました")
	}
	menue.Date = extractDateFromTitleString(head)

	doc.Find("div.productItem").Each(func(i int, productNode *goquery.Selection) {
		imageURL := productNode.Find("img").First().AttrOr("src", "")
		isOrdered := (productNode.Find("figure").First().AttrOr("class", "") != "")
		content := productNode.Find("div.productContent")
		maker := content.Find("span.truncate").Text()
		name := content.Find("p.truncate").Text()
		price := content.Find("div.productPrice").Text()
		price = extractPrice(price) + "円"
		amount := content.Find("div.productAmountInner").Find("input").AttrOr("value", "0")
		orderedAmount := 0
		if val, err := strconv.Atoi(amount); err == nil {
			orderedAmount = val
		}
		product := Product{
			Name:          name,
			Maker:         maker,
			ImageURL:      imageURL,
			IsOrdered:     isOrdered,
			Price:         price,
			OrderedAmount: orderedAmount,
		}
		menue.Products = append(menue.Products, product)
	})

	return menue, nil
}
