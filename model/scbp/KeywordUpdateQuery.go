package scbp

// KeywordUpdateQuery 结构体
type KeywordUpdateQuery struct {
	// 关键词集合
	KeywordList []KeywordInfo `json:"keyword_list,omitempty" xml:"keyword_list>keyword_info,omitempty"`
	// 更新类型：add_price（批量修改成不同价格，keywordList中price不能为空）update_price(批量修改成相同价格，updateInfo中price不能为空)
	UpdateType string `json:"update_type,omitempty" xml:"update_type,omitempty"`
	// 更新信息
	UpdateInfo *KeywordInfo `json:"update_info,omitempty" xml:"update_info,omitempty"`
}
