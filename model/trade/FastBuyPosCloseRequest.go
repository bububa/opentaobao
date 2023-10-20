package trade

import (
	"sync"
)

// FastBuyPosCloseRequest 结构体
type FastBuyPosCloseRequest struct {
	// 外部唯一订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 外部门店编码
	OutShopCode string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
}

var poolFastBuyPosCloseRequest = sync.Pool{
	New: func() any {
		return new(FastBuyPosCloseRequest)
	},
}

// GetFastBuyPosCloseRequest() 从对象池中获取FastBuyPosCloseRequest
func GetFastBuyPosCloseRequest() *FastBuyPosCloseRequest {
	return poolFastBuyPosCloseRequest.Get().(*FastBuyPosCloseRequest)
}

// ReleaseFastBuyPosCloseRequest 释放FastBuyPosCloseRequest
func ReleaseFastBuyPosCloseRequest(v *FastBuyPosCloseRequest) {
	v.OutOrderId = ""
	v.StoreId = ""
	v.OutShopCode = ""
	poolFastBuyPosCloseRequest.Put(v)
}
