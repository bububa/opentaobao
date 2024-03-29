package simba

import (
	"sync"
)

// BidwordDefaultQueryVo 结构体
type BidwordDefaultQueryVo struct {
	// 宝贝id集合
	MaterialIdList []int64 `json:"material_id_list,omitempty" xml:"material_id_list>int64,omitempty"`
}

var poolBidwordDefaultQueryVo = sync.Pool{
	New: func() any {
		return new(BidwordDefaultQueryVo)
	},
}

// GetBidwordDefaultQueryVo() 从对象池中获取BidwordDefaultQueryVo
func GetBidwordDefaultQueryVo() *BidwordDefaultQueryVo {
	return poolBidwordDefaultQueryVo.Get().(*BidwordDefaultQueryVo)
}

// ReleaseBidwordDefaultQueryVo 释放BidwordDefaultQueryVo
func ReleaseBidwordDefaultQueryVo(v *BidwordDefaultQueryVo) {
	v.MaterialIdList = v.MaterialIdList[:0]
	poolBidwordDefaultQueryVo.Put(v)
}
