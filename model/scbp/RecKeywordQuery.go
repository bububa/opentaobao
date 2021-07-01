package scbp

// RecKeywordQuery 结构体
type RecKeywordQuery struct {
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
	// 搜索词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}
