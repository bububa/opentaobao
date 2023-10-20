package tmallsc

import (
	"sync"
)

// ResultBase 结构体
type ResultBase struct {
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
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
	v.GmtModified = ""
	v.GmtCreate = ""
	v.Value = ""
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolResultBase.Put(v)
}
