package main

import (
	"fmt"

	"github.com/8245snake/gochikurunow"
)

func main() {
	//クライアントを初期化
	api, err := gochikurunow.NewGochiClient("your mail address", "your password")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//メニューを取得
	menue, err := api.GetMenu()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//情報を取得
	fmt.Printf("%s\n", menue.Date)
	for _, product := range menue.Products {
		fmt.Printf("%s\n", product.Maker)
		fmt.Printf("%s\n", product.Name)
		fmt.Printf("%s\n", product.ImageURL)
		fmt.Printf("%s\n", product.Price)
	}
}
