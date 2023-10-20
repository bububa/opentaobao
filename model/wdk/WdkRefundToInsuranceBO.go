package wdk

import (
	"sync"
)

// WdkRefundToInsuranceBO 结构体
type WdkRefundToInsuranceBO struct {
	// 退货原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 退款完成时间
	RefundSuccessTime string `json:"refund_success_time,omitempty" xml:"refund_success_time,omitempty"`
	// 收货地址
	DeliveryAddress string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	// 发货地址
	SendAddress string `json:"send_address,omitempty" xml:"send_address,omitempty"`
	// 退货金额
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 商品类目名称(从root到叶子节点)
	ItemCategory string `json:"item_category,omitempty" xml:"item_category,omitempty"`
	// 货物单价
	ItemPrice string `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 货物数量(下单销售数量)
	ItemQuantity string `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 货物名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 退货单ID
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款发起时间
	RefundCreateTime string `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 售后退=2，售中退=1
	ReverseType string `json:"reverse_type,omitempty" xml:"reverse_type,omitempty"`
	// 交易子订单ID
	TbSubOrderId int64 `json:"tb_sub_order_id,omitempty" xml:"tb_sub_order_id,omitempty"`
}

var poolWdkRefundToInsuranceBO = sync.Pool{
	New: func() any {
		return new(WdkRefundToInsuranceBO)
	},
}

// GetWdkRefundToInsuranceBO() 从对象池中获取WdkRefundToInsuranceBO
func GetWdkRefundToInsuranceBO() *WdkRefundToInsuranceBO {
	return poolWdkRefundToInsuranceBO.Get().(*WdkRefundToInsuranceBO)
}

// ReleaseWdkRefundToInsuranceBO 释放WdkRefundToInsuranceBO
func ReleaseWdkRefundToInsuranceBO(v *WdkRefundToInsuranceBO) {
	v.RefundReason = ""
	v.RefundSuccessTime = ""
	v.DeliveryAddress = ""
	v.SendAddress = ""
	v.RefundAmount = ""
	v.ItemCategory = ""
	v.ItemPrice = ""
	v.ItemQuantity = ""
	v.ItemName = ""
	v.RefundId = ""
	v.RefundCreateTime = ""
	v.SignTime = ""
	v.ReverseType = ""
	v.TbSubOrderId = 0
	poolWdkRefundToInsuranceBO.Put(v)
}
