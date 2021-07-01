package wangwang

// WordCountList 结构体
type WordCountList struct {
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 关键词数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
