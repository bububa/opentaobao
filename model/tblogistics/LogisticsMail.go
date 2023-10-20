package tblogistics

import (
	"sync"
)

// LogisticsMail 结构体
type LogisticsMail struct {
	// 运单号.具体一个物流公司的运单号码.
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
}

var poolLogisticsMail = sync.Pool{
	New: func() any {
		return new(LogisticsMail)
	},
}

// GetLogisticsMail() 从对象池中获取LogisticsMail
func GetLogisticsMail() *LogisticsMail {
	return poolLogisticsMail.Get().(*LogisticsMail)
}

// ReleaseLogisticsMail 释放LogisticsMail
func ReleaseLogisticsMail(v *LogisticsMail) {
	v.OutSid = ""
	v.CompanyName = ""
	poolLogisticsMail.Put(v)
}
