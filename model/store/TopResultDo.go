package store

import (
	"sync"
)

// TopResultDo 结构体
type TopResultDo struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果：true成功；false失败
	Result *FullStoreTopDto `json:"result,omitempty" xml:"result,omitempty"`
	// 个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResultDo = sync.Pool{
	New: func() any {
		return new(TopResultDo)
	},
}

// GetTopResultDo() 从对象池中获取TopResultDo
func GetTopResultDo() *TopResultDo {
	return poolTopResultDo.Get().(*TopResultDo)
}

// ReleaseTopResultDo 释放TopResultDo
func ReleaseTopResultDo(v *TopResultDo) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	v.TotalNum = 0
	v.Failure = false
	v.Success = false
	poolTopResultDo.Put(v)
}
