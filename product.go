package gochikurunow

//Product 商品
type Product struct {
	//名前
	Name string `json:"name"`
	//販売元
	Maker string `json:"maker"`
	//金額（◯◯◯円）
	Price string `json:"price"`
	//写真のURL
	ImageURL string `json:"image_url"`
	//自分が注文したか
	IsOrdered bool `json:"is_ordered"`
	//注文数
	OrderedAmount int `json:"amount"`
}

//MenueInfo メニュー
type MenueInfo struct {
	//対象の日付（◯月◯日（曜日））
	Date string `json:"date"`
	//選択できる商品
	Products []Product `json:"products"`
}
