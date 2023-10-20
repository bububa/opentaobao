package simba

import (
	"sync"
)

// BidStrategyVo 结构体
type BidStrategyVo struct {
	// 名称
	RankName string `json:"rank_name,omitempty" xml:"rank_name,omitempty"`
	// 溢价
	Dicount int64 `json:"dicount,omitempty" xml:"dicount,omitempty"`
	// 0停止，1生效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 抢位排名
	Rank int64 `json:"rank,omitempty" xml:"rank,omitempty"`
}

var poolBidStrategyVo = sync.Pool{
	New: func() any {
		return new(BidStrategyVo)
	},
}

// GetBidStrategyVo() 从对象池中获取BidStrategyVo
func GetBidStrategyVo() *BidStrategyVo {
	return poolBidStrategyVo.Get().(*BidStrategyVo)
}

// ReleaseBidStrategyVo 释放BidStrategyVo
func ReleaseBidStrategyVo(v *BidStrategyVo) {
	v.RankName = ""
	v.Dicount = 0
	v.Status = 0
	v.Rank = 0
	poolBidStrategyVo.Put(v)
}
