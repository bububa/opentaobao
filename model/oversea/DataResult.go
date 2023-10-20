package oversea

import (
	"sync"
)

// DataResult 结构体
type DataResult struct {
	// 查到的税率信息
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 错误代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 查询结果是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataResult = sync.Pool{
	New: func() any {
		return new(DataResult)
	},
}

// GetDataResult() 从对象池中获取DataResult
func GetDataResult() *DataResult {
	return poolDataResult.Get().(*DataResult)
}

// ReleaseDataResult 释放DataResult
func ReleaseDataResult(v *DataResult) {
	v.ExchangeRate = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolDataResult.Put(v)
}
