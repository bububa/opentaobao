package openmall

import (
	"sync"
)

// TopTradeResultVo 结构体
type TopTradeResultVo struct {
	// 运费列表
	Posts []PostDo `json:"posts,omitempty" xml:"posts>post_do,omitempty"`
	// 淘宝交易ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 发货地址
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 发货地址对应的areaid
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolTopTradeResultVo = sync.Pool{
	New: func() any {
		return new(TopTradeResultVo)
	},
}

// GetTopTradeResultVo() 从对象池中获取TopTradeResultVo
func GetTopTradeResultVo() *TopTradeResultVo {
	return poolTopTradeResultVo.Get().(*TopTradeResultVo)
}

// ReleaseTopTradeResultVo 释放TopTradeResultVo
func ReleaseTopTradeResultVo(v *TopTradeResultVo) {
	v.Posts = v.Posts[:0]
	v.Tid = ""
	v.Location = ""
	v.AreaId = 0
	v.ItemId = 0
	poolTopTradeResultVo.Put(v)
}
