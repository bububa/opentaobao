package tmallcar

import (
	"sync"
)

// ConsumeEticketCommand 结构体
type ConsumeEticketCommand struct {
	// 核销码
	EticketCode string `json:"eticket_code,omitempty" xml:"eticket_code,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolConsumeEticketCommand = sync.Pool{
	New: func() any {
		return new(ConsumeEticketCommand)
	},
}

// GetConsumeEticketCommand() 从对象池中获取ConsumeEticketCommand
func GetConsumeEticketCommand() *ConsumeEticketCommand {
	return poolConsumeEticketCommand.Get().(*ConsumeEticketCommand)
}

// ReleaseConsumeEticketCommand 释放ConsumeEticketCommand
func ReleaseConsumeEticketCommand(v *ConsumeEticketCommand) {
	v.EticketCode = ""
	v.OrderId = 0
	v.StoreId = 0
	poolConsumeEticketCommand.Put(v)
}
