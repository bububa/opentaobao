package trade

import (
	"sync"
)

// IsvCreateOrderParam 结构体
type IsvCreateOrderParam struct {
	// 订单行详情
	SubOrderInfoList []IsvSimpleSubOrderModel `json:"sub_order_info_list,omitempty" xml:"sub_order_info_list>isv_simple_sub_order_model,omitempty"`
	// 买家备注
	BuyerRemarks string `json:"buyer_remarks,omitempty" xml:"buyer_remarks,omitempty"`
	// 外部订单ID，仅仅做追踪使用
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 收货地址信息
	SmAddrModel *SmAddrModel `json:"sm_addr_model,omitempty" xml:"sm_addr_model,omitempty"`
}

var poolIsvCreateOrderParam = sync.Pool{
	New: func() any {
		return new(IsvCreateOrderParam)
	},
}

// GetIsvCreateOrderParam() 从对象池中获取IsvCreateOrderParam
func GetIsvCreateOrderParam() *IsvCreateOrderParam {
	return poolIsvCreateOrderParam.Get().(*IsvCreateOrderParam)
}

// ReleaseIsvCreateOrderParam 释放IsvCreateOrderParam
func ReleaseIsvCreateOrderParam(v *IsvCreateOrderParam) {
	v.SubOrderInfoList = v.SubOrderInfoList[:0]
	v.BuyerRemarks = ""
	v.OutOrderId = ""
	v.SmAddrModel = nil
	poolIsvCreateOrderParam.Put(v)
}
