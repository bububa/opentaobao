package tmallcar

// TradeItemDto 结构体
type TradeItemDto struct {
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 商品主图片URL
	ItemPictUrl string `json:"item_pict_url,omitempty" xml:"item_pict_url,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
