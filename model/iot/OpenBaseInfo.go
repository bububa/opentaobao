package iot

import (
	"sync"
)

// OpenBaseInfo 结构体
type OpenBaseInfo struct {
	// 账户体系隔离
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
	// 用户ID，此处传入第三方账户体系的用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	UtdId string `json:"utd_id,omitempty" xml:"utd_id,omitempty"`
	// 扩展信息，用于存放APP类型等
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
}

var poolOpenBaseInfo = sync.Pool{
	New: func() any {
		return new(OpenBaseInfo)
	},
}

// GetOpenBaseInfo() 从对象池中获取OpenBaseInfo
func GetOpenBaseInfo() *OpenBaseInfo {
	return poolOpenBaseInfo.Get().(*OpenBaseInfo)
}

// ReleaseOpenBaseInfo 释放OpenBaseInfo
func ReleaseOpenBaseInfo(v *OpenBaseInfo) {
	v.Schema = ""
	v.UserId = ""
	v.UtdId = ""
	v.Ext = ""
	poolOpenBaseInfo.Put(v)
}
