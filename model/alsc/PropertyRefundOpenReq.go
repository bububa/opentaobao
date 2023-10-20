package alsc

import (
	"sync"
)

// PropertyRefundOpenReq 结构体
type PropertyRefundOpenReq struct {
	// 券实例id集合
	VoucherIdList []string `json:"voucher_id_list,omitempty" xml:"voucher_id_list>string,omitempty"`
	// saas品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 退款的订单号
	NewOuterOrderId string `json:"new_outer_order_id,omitempty" xml:"new_outer_order_id,omitempty"`
	// 操作人
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 原订单号，就是核销的订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 外部id类型，wechat微信，alipay支付宝
	OuterType string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 请求幂等id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
}

var poolPropertyRefundOpenReq = sync.Pool{
	New: func() any {
		return new(PropertyRefundOpenReq)
	},
}

// GetPropertyRefundOpenReq() 从对象池中获取PropertyRefundOpenReq
func GetPropertyRefundOpenReq() *PropertyRefundOpenReq {
	return poolPropertyRefundOpenReq.Get().(*PropertyRefundOpenReq)
}

// ReleasePropertyRefundOpenReq 释放PropertyRefundOpenReq
func ReleasePropertyRefundOpenReq(v *PropertyRefundOpenReq) {
	v.VoucherIdList = v.VoucherIdList[:0]
	v.BrandId = ""
	v.Mobile = ""
	v.NewOuterOrderId = ""
	v.OperatorId = ""
	v.OuterId = ""
	v.OuterOrderId = ""
	v.OuterType = ""
	v.RequestId = ""
	v.ShopId = ""
	v.CustomerId = ""
	poolPropertyRefundOpenReq.Put(v)
}
