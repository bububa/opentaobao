package eticket

import (
	"sync"
)

// EticketOpLog 结构体
type EticketOpLog struct {
	// 操作流水号
	ConsumeSerialNum string `json:"consume_serial_num,omitempty" xml:"consume_serial_num,omitempty"`
	// 操作金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 操作时间
	OpTime string `json:"op_time,omitempty" xml:"op_time,omitempty"`
	// 操作员身份ID
	PosId string `json:"pos_id,omitempty" xml:"pos_id,omitempty"`
	// 手机号码后四位
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 操作数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 操作类型 1:核销 2:冲正
	OpType int64 `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolEticketOpLog = sync.Pool{
	New: func() any {
		return new(EticketOpLog)
	},
}

// GetEticketOpLog() 从对象池中获取EticketOpLog
func GetEticketOpLog() *EticketOpLog {
	return poolEticketOpLog.Get().(*EticketOpLog)
}

// ReleaseEticketOpLog 释放EticketOpLog
func ReleaseEticketOpLog(v *EticketOpLog) {
	v.ConsumeSerialNum = ""
	v.Amount = ""
	v.OpTime = ""
	v.PosId = ""
	v.Mobile = ""
	v.Num = 0
	v.OpType = 0
	v.OrderId = 0
	poolEticketOpLog.Put(v)
}
