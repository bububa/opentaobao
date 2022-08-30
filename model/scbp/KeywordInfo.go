package scbp

// KeywordInfo 结构体
type KeywordInfo struct {
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 关键词id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}
