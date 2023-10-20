package logistic

import (
	"sync"
)

// Resultdata 结构体
type Resultdata struct {
	// 主交易单号
	ParentOrderId int64 `json:"parent_order_id,omitempty" xml:"parent_order_id,omitempty"`
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
	v.ParentOrderId = 0
	poolResultdata.Put(v)
}
