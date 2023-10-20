package tmc

import (
	"sync"
)

// TmcQueueInfo 结构体
type TmcQueueInfo struct {
	// TMC组名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 消息队列Broker名称
	BrokerName string `json:"broker_name,omitempty" xml:"broker_name,omitempty"`
	// 当前队列当天读取量
	GetTotal int64 `json:"get_total,omitempty" xml:"get_total,omitempty"`
	// 当前队列当天写入量
	PutToal int64 `json:"put_toal,omitempty" xml:"put_toal,omitempty"`
}

var poolTmcQueueInfo = sync.Pool{
	New: func() any {
		return new(TmcQueueInfo)
	},
}

// GetTmcQueueInfo() 从对象池中获取TmcQueueInfo
func GetTmcQueueInfo() *TmcQueueInfo {
	return poolTmcQueueInfo.Get().(*TmcQueueInfo)
}

// ReleaseTmcQueueInfo 释放TmcQueueInfo
func ReleaseTmcQueueInfo(v *TmcQueueInfo) {
	v.Name = ""
	v.BrokerName = ""
	v.GetTotal = 0
	v.PutToal = 0
	poolTmcQueueInfo.Put(v)
}
