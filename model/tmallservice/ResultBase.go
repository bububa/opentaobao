package tmallservice

import (
	"sync"
)

// ResultBase 结构体
type ResultBase struct {
	// 任务id列表
	Values []int64 `json:"values,omitempty" xml:"values>int64,omitempty"`
	// 服务预警消息列表
	ValueList []ServiceMonitorMessage `json:"value_list,omitempty" xml:"value_list>service_monitor_message,omitempty"`
	// 单个一键求助单对象json字符串
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultBase = sync.Pool{
	New: func() any {
		return new(ResultBase)
	},
}

// GetResultBase() 从对象池中获取ResultBase
func GetResultBase() *ResultBase {
	return poolResultBase.Get().(*ResultBase)
}

// ReleaseResultBase 释放ResultBase
func ReleaseResultBase(v *ResultBase) {
	v.Values = v.Values[:0]
	v.ValueList = v.ValueList[:0]
	v.Value = ""
	v.ErrorMsg = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.ErrorCode = 0
	v.Success = false
	poolResultBase.Put(v)
}
