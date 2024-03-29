package ascpchannel

import (
	"sync"
)

// ExtSubOrderCheckRequest 结构体
type ExtSubOrderCheckRequest struct {
	// 产品id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 外部子订单号
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolExtSubOrderCheckRequest = sync.Pool{
	New: func() any {
		return new(ExtSubOrderCheckRequest)
	},
}

// GetExtSubOrderCheckRequest() 从对象池中获取ExtSubOrderCheckRequest
func GetExtSubOrderCheckRequest() *ExtSubOrderCheckRequest {
	return poolExtSubOrderCheckRequest.Get().(*ExtSubOrderCheckRequest)
}

// ReleaseExtSubOrderCheckRequest 释放ExtSubOrderCheckRequest
func ReleaseExtSubOrderCheckRequest(v *ExtSubOrderCheckRequest) {
	v.ProductId = ""
	v.OutSubOrderNo = ""
	v.Quantity = 0
	poolExtSubOrderCheckRequest.Put(v)
}
