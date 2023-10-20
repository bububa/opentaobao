package store

import (
	"sync"
)

// UpdateResultDo 结构体
type UpdateResultDo struct {
	// 失败的标签id集合
	FailuredList []int64 `json:"failured_list,omitempty" xml:"failured_list>int64,omitempty"`
	// 成功的标签id集合
	SuccessList []int64 `json:"success_list,omitempty" xml:"success_list>int64,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 完整错误信息
	FullErrorMsg string `json:"full_error_msg,omitempty" xml:"full_error_msg,omitempty"`
	// 关键主键
	PriKey string `json:"pri_key,omitempty" xml:"pri_key,omitempty"`
	// 结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 成功数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 模型
	Models *Models `json:"models,omitempty" xml:"models,omitempty"`
	// 请求更新的总数量(已废弃)
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否失败
	Failured bool `json:"failured,omitempty" xml:"failured,omitempty"`
	// 是否请求成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUpdateResultDo = sync.Pool{
	New: func() any {
		return new(UpdateResultDo)
	},
}

// GetUpdateResultDo() 从对象池中获取UpdateResultDo
func GetUpdateResultDo() *UpdateResultDo {
	return poolUpdateResultDo.Get().(*UpdateResultDo)
}

// ReleaseUpdateResultDo 释放UpdateResultDo
func ReleaseUpdateResultDo(v *UpdateResultDo) {
	v.FailuredList = v.FailuredList[:0]
	v.SuccessList = v.SuccessList[:0]
	v.ErrorMsg = ""
	v.FullErrorMsg = ""
	v.PriKey = ""
	v.Result = ""
	v.ResultCode = ""
	v.Count = 0
	v.Models = nil
	v.TotalNum = 0
	v.Failured = false
	v.Success = false
	poolUpdateResultDo.Put(v)
}
