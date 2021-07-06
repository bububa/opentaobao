package c2m

// ItemTopDto 结构体
type ItemTopDto struct {
	// 主图信息
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户的nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 商品对应的类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 用户的id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
