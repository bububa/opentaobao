package trade

import (
	"sync"
)

// FastBuyPosQueryRequest 结构体
type FastBuyPosQueryRequest struct {
	// pos机id
	MachineId string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
	// 外部订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 外部门店编码
	OutShopCode string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
}

var poolFastBuyPosQueryRequest = sync.Pool{
	New: func() any {
		return new(FastBuyPosQueryRequest)
	},
}

// GetFastBuyPosQueryRequest() 从对象池中获取FastBuyPosQueryRequest
func GetFastBuyPosQueryRequest() *FastBuyPosQueryRequest {
	return poolFastBuyPosQueryRequest.Get().(*FastBuyPosQueryRequest)
}

// ReleaseFastBuyPosQueryRequest 释放FastBuyPosQueryRequest
func ReleaseFastBuyPosQueryRequest(v *FastBuyPosQueryRequest) {
	v.MachineId = ""
	v.OutOrderId = ""
	v.StoreId = ""
	v.OutShopCode = ""
	poolFastBuyPosQueryRequest.Put(v)
}
