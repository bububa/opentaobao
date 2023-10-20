package simba

import (
	"sync"
)

// RankedItem 结构体
type RankedItem struct {
	// 客户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 最高出价
	MaxPrice string `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 创意标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 排名
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 原始质量评分
	RankScore int64 `json:"rank_score,omitempty" xml:"rank_score,omitempty"`
}

var poolRankedItem = sync.Pool{
	New: func() any {
		return new(RankedItem)
	},
}

// GetRankedItem() 从对象池中获取RankedItem
func GetRankedItem() *RankedItem {
	return poolRankedItem.Get().(*RankedItem)
}

// ReleaseRankedItem 释放RankedItem
func ReleaseRankedItem(v *RankedItem) {
	v.Nick = ""
	v.MaxPrice = ""
	v.Title = ""
	v.LinkUrl = ""
	v.Order = 0
	v.RankScore = 0
	poolRankedItem.Put(v)
}
