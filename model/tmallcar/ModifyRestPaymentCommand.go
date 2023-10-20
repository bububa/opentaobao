package tmallcar

import (
	"sync"
)

// ModifyRestPaymentCommand 结构体
type ModifyRestPaymentCommand struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 期望改到的实付金额，单位分
	TargetActualPayFee int64 `json:"target_actual_pay_fee,omitempty" xml:"target_actual_pay_fee,omitempty"`
	// 子支付单id
	SubPayOrderId int64 `json:"sub_pay_order_id,omitempty" xml:"sub_pay_order_id,omitempty"`
}

var poolModifyRestPaymentCommand = sync.Pool{
	New: func() any {
		return new(ModifyRestPaymentCommand)
	},
}

// GetModifyRestPaymentCommand() 从对象池中获取ModifyRestPaymentCommand
func GetModifyRestPaymentCommand() *ModifyRestPaymentCommand {
	return poolModifyRestPaymentCommand.Get().(*ModifyRestPaymentCommand)
}

// ReleaseModifyRestPaymentCommand 释放ModifyRestPaymentCommand
func ReleaseModifyRestPaymentCommand(v *ModifyRestPaymentCommand) {
	v.OrderId = 0
	v.TargetActualPayFee = 0
	v.SubPayOrderId = 0
	poolModifyRestPaymentCommand.Put(v)
}
