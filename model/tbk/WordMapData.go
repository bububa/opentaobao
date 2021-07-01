package tbk

// WordMapData 结构体
type WordMapData struct {
	// 链接-商品相关关联词落地页地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 商品相关的关联词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
}
