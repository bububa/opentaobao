package bus

import (
	"sync"
)

// BusoMainOrderHistoryPageVo 结构体
type BusoMainOrderHistoryPageVo struct {
	// busoMainOrderHistoryVOList
	BusoMainOrderHistoryVOList []Busomainorderhistoryvolist `json:"buso_main_order_history_v_o_list,omitempty" xml:"buso_main_order_history_v_o_list>busomainorderhistoryvolist,omitempty"`
	// currentPage 当前多少页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// totalSize 总条目
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

var poolBusoMainOrderHistoryPageVo = sync.Pool{
	New: func() any {
		return new(BusoMainOrderHistoryPageVo)
	},
}

// GetBusoMainOrderHistoryPageVo() 从对象池中获取BusoMainOrderHistoryPageVo
func GetBusoMainOrderHistoryPageVo() *BusoMainOrderHistoryPageVo {
	return poolBusoMainOrderHistoryPageVo.Get().(*BusoMainOrderHistoryPageVo)
}

// ReleaseBusoMainOrderHistoryPageVo 释放BusoMainOrderHistoryPageVo
func ReleaseBusoMainOrderHistoryPageVo(v *BusoMainOrderHistoryPageVo) {
	v.BusoMainOrderHistoryVOList = v.BusoMainOrderHistoryVOList[:0]
	v.CurrentPage = 0
	v.TotalSize = 0
	poolBusoMainOrderHistoryPageVo.Put(v)
}
