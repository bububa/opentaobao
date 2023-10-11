package simba

// SuggestBidwordVo 结构体
type SuggestBidwordVo struct {
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 关键词
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 市场平均出价
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 预估展现
	Impression string `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 点击转化率
	Cvr string `json:"cvr,omitempty" xml:"cvr,omitempty"`
	// 相关性,1:差,2:中,3:好
	RelevanceType int64 `json:"relevance_type,omitempty" xml:"relevance_type,omitempty"`
	// 分类,0:宝贝,1:店铺
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
