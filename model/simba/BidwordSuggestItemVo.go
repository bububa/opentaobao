package simba

import (
	"sync"
)

// BidwordSuggestItemVo 结构体
type BidwordSuggestItemVo struct {
	// 关键词建议条目集合
	WordList []SuggestBidwordVo `json:"word_list,omitempty" xml:"word_list>suggest_bidword_vo,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}

var poolBidwordSuggestItemVo = sync.Pool{
	New: func() any {
		return new(BidwordSuggestItemVo)
	},
}

// GetBidwordSuggestItemVo() 从对象池中获取BidwordSuggestItemVo
func GetBidwordSuggestItemVo() *BidwordSuggestItemVo {
	return poolBidwordSuggestItemVo.Get().(*BidwordSuggestItemVo)
}

// ReleaseBidwordSuggestItemVo 释放BidwordSuggestItemVo
func ReleaseBidwordSuggestItemVo(v *BidwordSuggestItemVo) {
	v.WordList = v.WordList[:0]
	v.MaterialId = 0
	poolBidwordSuggestItemVo.Put(v)
}
