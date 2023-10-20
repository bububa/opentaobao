package util

import (
	"sync"
)

// QimenEvent 结构体
type QimenEvent struct {
	// 奇门事件对象
	Event *Event `json:"event,omitempty" xml:"event,omitempty"`
}

var poolQimenEvent = sync.Pool{
	New: func() any {
		return new(QimenEvent)
	},
}

// GetQimenEvent() 从对象池中获取QimenEvent
func GetQimenEvent() *QimenEvent {
	return poolQimenEvent.Get().(*QimenEvent)
}

// ReleaseQimenEvent 释放QimenEvent
func ReleaseQimenEvent(v *QimenEvent) {
	v.Event = nil
	poolQimenEvent.Put(v)
}
