package tbk

// NTbkShop 结构体
type NTbkShop struct {
	// 店铺推广链接
	Clickurl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 店铺地址
	Shopurl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 卖家昵称
	Sellernick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 店铺类型，B：天猫，C：淘宝
	Shoptype string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 店铺名称
	Shoptitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 店标图片
	Picturl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 卖家id
	Userid int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
