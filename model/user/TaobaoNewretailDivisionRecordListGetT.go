package user

import (
	"sync"
)

// TaobaoNewretailDivisionRecordListGetT 结构体
type TaobaoNewretailDivisionRecordListGetT struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 导购昵称
	CommissionEmployeeName string `json:"commission_employee_name,omitempty" xml:"commission_employee_name,omitempty"`
	// 订单创建时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 对应账号的OpenUID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 订单结束时间
	OrderEndTime string `json:"order_end_time,omitempty" xml:"order_end_time,omitempty"`
	// 分佣金额
	CommissionMoney int64 `json:"commission_money,omitempty" xml:"commission_money,omitempty"`
	// 主订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单状态 12-付款，13-关闭，14-确认收货，3-结算成功
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 分佣时间
	CommissionTime int64 `json:"commission_time,omitempty" xml:"commission_time,omitempty"`
	// 子订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 导购id
	CommissionEmployeeId int64 `json:"commission_employee_id,omitempty" xml:"commission_employee_id,omitempty"`
	// 订单支付时间
	OrderPayTime int64 `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
	// 订单支付金额
	OrderPayMoney int64 `json:"order_pay_money,omitempty" xml:"order_pay_money,omitempty"`
}

var poolTaobaoNewretailDivisionRecordListGetT = sync.Pool{
	New: func() any {
		return new(TaobaoNewretailDivisionRecordListGetT)
	},
}

// GetTaobaoNewretailDivisionRecordListGetT() 从对象池中获取TaobaoNewretailDivisionRecordListGetT
func GetTaobaoNewretailDivisionRecordListGetT() *TaobaoNewretailDivisionRecordListGetT {
	return poolTaobaoNewretailDivisionRecordListGetT.Get().(*TaobaoNewretailDivisionRecordListGetT)
}

// ReleaseTaobaoNewretailDivisionRecordListGetT 释放TaobaoNewretailDivisionRecordListGetT
func ReleaseTaobaoNewretailDivisionRecordListGetT(v *TaobaoNewretailDivisionRecordListGetT) {
	v.BuyerNick = ""
	v.CommissionEmployeeName = ""
	v.OrderCreateTime = ""
	v.OpenUid = ""
	v.OrderEndTime = ""
	v.CommissionMoney = 0
	v.OrderId = 0
	v.OrderStatus = 0
	v.CommissionTime = 0
	v.BizOrderId = 0
	v.CommissionEmployeeId = 0
	v.OrderPayTime = 0
	v.OrderPayMoney = 0
	poolTaobaoNewretailDivisionRecordListGetT.Put(v)
}
