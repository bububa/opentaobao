package simba

// ShopInfoVo 结构体
type ShopInfoVo struct {
	// 店铺名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}
