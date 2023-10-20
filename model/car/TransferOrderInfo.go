package car

import (
	"sync"
)

// TransferOrderInfo 结构体
type TransferOrderInfo struct {
	// 订单修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 外部商家订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 订单创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 打款时间
	PaymentTime string `json:"payment_time,omitempty" xml:"payment_time,omitempty"`
	// 关单/取消 原因(没有时为空)
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 退款时间
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// 出发地至目的地之间距离 单位(米)
	Distance int64 `json:"distance,omitempty" xml:"distance,omitempty"`
	// 订单状态   21-等待商家确认接单，22-商家已确认接单（未派司机），23-商家已确认接单（已派司机），24-司机服务已完成，25-司机已出发，26-司机已到达目的地，27-司机开始服务 ，60-订单已关闭，70-订单已完成
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 是否有责取消（可能为空）
	CloseDuty bool `json:"close_duty,omitempty" xml:"close_duty,omitempty"`
}

var poolTransferOrderInfo = sync.Pool{
	New: func() any {
		return new(TransferOrderInfo)
	},
}

// GetTransferOrderInfo() 从对象池中获取TransferOrderInfo
func GetTransferOrderInfo() *TransferOrderInfo {
	return poolTransferOrderInfo.Get().(*TransferOrderInfo)
}

// ReleaseTransferOrderInfo 释放TransferOrderInfo
func ReleaseTransferOrderInfo(v *TransferOrderInfo) {
	v.ModifiedTime = ""
	v.OrderId = ""
	v.OutOrderId = ""
	v.CreatedTime = ""
	v.PaymentTime = ""
	v.CancelReason = ""
	v.RefundTime = ""
	v.Distance = 0
	v.OrderStatus = 0
	v.CloseDuty = false
	poolTransferOrderInfo.Put(v)
}
