package wdk

import (
	"sync"
)

// OrderSuccessResponse 结构体
type OrderSuccessResponse struct {
	// 子单列表
	OrderSubInfoList []OrderSubInfoBO `json:"order_sub_info_list,omitempty" xml:"order_sub_info_list>order_sub_info_bo,omitempty"`
	// 返回的业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolOrderSuccessResponse = sync.Pool{
	New: func() any {
		return new(OrderSuccessResponse)
	},
}

// GetOrderSuccessResponse() 从对象池中获取OrderSuccessResponse
func GetOrderSuccessResponse() *OrderSuccessResponse {
	return poolOrderSuccessResponse.Get().(*OrderSuccessResponse)
}

// ReleaseOrderSuccessResponse 释放OrderSuccessResponse
func ReleaseOrderSuccessResponse(v *OrderSuccessResponse) {
	v.OrderSubInfoList = v.OrderSubInfoList[:0]
	v.BizOrderId = ""
	v.OutOrderId = ""
	poolOrderSuccessResponse.Put(v)
}
