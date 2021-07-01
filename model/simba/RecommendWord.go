package simba

// RecommendWord 结构体
type RecommendWord struct {
	// 推荐的关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 搜索量
	Pv string `json:"pv,omitempty" xml:"pv,omitempty"`
	// 平均价格
	AveragePrice string `json:"average_price,omitempty" xml:"average_price,omitempty"`
	// 相关度
	Pertinence string `json:"pertinence,omitempty" xml:"pertinence,omitempty"`
}
