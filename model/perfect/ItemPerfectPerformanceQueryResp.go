package perfect

import (
	"sync"
)

// ItemPerfectPerformanceQueryResp 结构体
type ItemPerfectPerformanceQueryResp struct {
	// 完美履约门店列表
	TcStoreList []TcStoreList `json:"tc_store_list,omitempty" xml:"tc_store_list>tc_store_list,omitempty"`
	// 商家商品id
	ItemOuterId string `json:"item_outer_id,omitempty" xml:"item_outer_id,omitempty"`
	// 天猫商品id
	ItemTmallId string `json:"item_tmall_id,omitempty" xml:"item_tmall_id,omitempty"`
	// 是否完美履约商品
	IsTcItem bool `json:"is_tc_item,omitempty" xml:"is_tc_item,omitempty"`
}

var poolItemPerfectPerformanceQueryResp = sync.Pool{
	New: func() any {
		return new(ItemPerfectPerformanceQueryResp)
	},
}

// GetItemPerfectPerformanceQueryResp() 从对象池中获取ItemPerfectPerformanceQueryResp
func GetItemPerfectPerformanceQueryResp() *ItemPerfectPerformanceQueryResp {
	return poolItemPerfectPerformanceQueryResp.Get().(*ItemPerfectPerformanceQueryResp)
}

// ReleaseItemPerfectPerformanceQueryResp 释放ItemPerfectPerformanceQueryResp
func ReleaseItemPerfectPerformanceQueryResp(v *ItemPerfectPerformanceQueryResp) {
	v.TcStoreList = v.TcStoreList[:0]
	v.ItemOuterId = ""
	v.ItemTmallId = ""
	v.IsTcItem = false
	poolItemPerfectPerformanceQueryResp.Put(v)
}
