package product

// GoodsSearchResult 结构体
type GoodsSearchResult struct {
	// 搜素结果所在的类目信息
	Category *WholesaleCategory `json:"category,omitempty" xml:"category,omitempty"`
	// 返回结果数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 搜索产品列表
	Items []GoodsSummary `json:"items,omitempty" xml:"items>goods_summary,omitempty"`
}
