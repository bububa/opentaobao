package alihealthpw

import (
	"sync"
)

// WriterInfo 结构体
type WriterInfo struct {
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 联系方式
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收入
	Income string `json:"income,omitempty" xml:"income,omitempty"`
	// 亲属关系
	RelationShip string `json:"relation_ship,omitempty" xml:"relation_ship,omitempty"`
}

var poolWriterInfo = sync.Pool{
	New: func() any {
		return new(WriterInfo)
	},
}

// GetWriterInfo() 从对象池中获取WriterInfo
func GetWriterInfo() *WriterInfo {
	return poolWriterInfo.Get().(*WriterInfo)
}

// ReleaseWriterInfo 释放WriterInfo
func ReleaseWriterInfo(v *WriterInfo) {
	v.Channel = ""
	v.Phone = ""
	v.Income = ""
	v.RelationShip = ""
	poolWriterInfo.Put(v)
}
