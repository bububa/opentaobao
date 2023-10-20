package refund

import (
	"sync"
)

// Resultdata 结构体
type Resultdata struct {
	// 数据消费结果编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 数据消费结果提示信息
	ResultTips string `json:"result_tips,omitempty" xml:"result_tips,omitempty"`
	// 数据消费状态
	ConsumeStatus string `json:"consume_status,omitempty" xml:"consume_status,omitempty"`
	// 退款单号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
}

var poolResultdata = sync.Pool{
	New: func() any {
		return new(Resultdata)
	},
}

// GetResultdata() 从对象池中获取Resultdata
func GetResultdata() *Resultdata {
	return poolResultdata.Get().(*Resultdata)
}

// ReleaseResultdata 释放Resultdata
func ReleaseResultdata(v *Resultdata) {
	v.ResultCode = ""
	v.ResultTips = ""
	v.ConsumeStatus = ""
	v.RefundId = 0
	poolResultdata.Put(v)
}
