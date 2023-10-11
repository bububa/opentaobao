package simba

// BidwordSuggestItemVo 结构体
type BidwordSuggestItemVo struct {
	// 关键词建议条目集合
	WordList []SuggestBidwordVo `json:"word_list,omitempty" xml:"word_list>suggest_bidword_vo,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}
