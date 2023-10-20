package westcrm

import (
	"sync"
)

// DataResult 结构体
type DataResult struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	DataList *CustomerBaseInfoVo `json:"data_list,omitempty" xml:"data_list,omitempty"`
	// 参数code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *UserConsumerVo `json:"data,omitempty" xml:"data,omitempty"`
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
	v.Message = ""
	v.DataList = nil
	v.Code = 0
	v.Data = nil
	poolDataResult.Put(v)
}
