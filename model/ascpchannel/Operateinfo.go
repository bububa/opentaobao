package ascpchannel

import (
	"sync"
)

// Operateinfo 结构体
type Operateinfo struct {
	// 操作时间Date类型
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 操作人
	OperateName string `json:"operate_name,omitempty" xml:"operate_name,omitempty"`
}

var poolOperateinfo = sync.Pool{
	New: func() any {
		return new(Operateinfo)
	},
}

// GetOperateinfo() 从对象池中获取Operateinfo
func GetOperateinfo() *Operateinfo {
	return poolOperateinfo.Get().(*Operateinfo)
}

// ReleaseOperateinfo 释放Operateinfo
func ReleaseOperateinfo(v *Operateinfo) {
	v.OperateTime = ""
	v.OperateName = ""
	poolOperateinfo.Put(v)
}
