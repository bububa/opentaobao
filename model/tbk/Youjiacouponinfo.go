package tbk

// Youjiacouponinfo 结构体
type Youjiacouponinfo struct {
	// 有价券商品id
	Itemid string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
