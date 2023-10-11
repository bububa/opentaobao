package idleisv

// AppraiseIsvItemDto 结构体
type AppraiseIsvItemDto struct {
	// 商品图片,绝对途径
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品Id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品价格，单位分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
