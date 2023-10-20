package tbk

// RecommendItemList 结构体
type RecommendItemList struct {
	// 商品链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 权益推荐商品id
	Itemid int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
