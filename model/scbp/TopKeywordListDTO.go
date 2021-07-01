package scbp

// TopKeywordListDto 结构体
type TopKeywordListDto struct {
	// 关键词列表
	KeywordList []string `json:"keyword_list,omitempty" xml:"keyword_list>string,omitempty"`
}
