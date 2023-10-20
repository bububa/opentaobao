package store

import (
	"sync"
)

// TopBatchResultDo 结构体
type TopBatchResultDo struct {
	// 失败的门店id
	ResultList []int64 `json:"result_list,omitempty" xml:"result_list>int64,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 其它信息
	Other *Other `json:"other,omitempty" xml:"other,omitempty"`
	// 总条目数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopBatchResultDo = sync.Pool{
	New: func() any {
		return new(TopBatchResultDo)
	},
}

// GetTopBatchResultDo() 从对象池中获取TopBatchResultDo
func GetTopBatchResultDo() *TopBatchResultDo {
	return poolTopBatchResultDo.Get().(*TopBatchResultDo)
}

// ReleaseTopBatchResultDo 释放TopBatchResultDo
func ReleaseTopBatchResultDo(v *TopBatchResultDo) {
	v.ResultList = v.ResultList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Other = nil
	v.TotalNum = 0
	v.Failure = false
	v.Success = false
	poolTopBatchResultDo.Put(v)
}
