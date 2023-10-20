package alidoc

import (
	"sync"
)

// FailedList 结构体
type FailedList struct {
	// 获取处方失败的订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 失败信息
	FailedMessage string `json:"failed_message,omitempty" xml:"failed_message,omitempty"`
}

var poolFailedList = sync.Pool{
	New: func() any {
		return new(FailedList)
	},
}

// GetFailedList() 从对象池中获取FailedList
func GetFailedList() *FailedList {
	return poolFailedList.Get().(*FailedList)
}

// ReleaseFailedList 释放FailedList
func ReleaseFailedList(v *FailedList) {
	v.BizOrderId = ""
	v.FailedMessage = ""
	poolFailedList.Put(v)
}
