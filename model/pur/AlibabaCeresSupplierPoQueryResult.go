package pur

import (
	"sync"
)

// AlibabaCeresSupplierPoQueryResult 结构体
type AlibabaCeresSupplierPoQueryResult struct {
	// 返回单据消息体List
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否查询成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaCeresSupplierPoQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaCeresSupplierPoQueryResult)
	},
}

// GetAlibabaCeresSupplierPoQueryResult() 从对象池中获取AlibabaCeresSupplierPoQueryResult
func GetAlibabaCeresSupplierPoQueryResult() *AlibabaCeresSupplierPoQueryResult {
	return poolAlibabaCeresSupplierPoQueryResult.Get().(*AlibabaCeresSupplierPoQueryResult)
}

// ReleaseAlibabaCeresSupplierPoQueryResult 释放AlibabaCeresSupplierPoQueryResult
func ReleaseAlibabaCeresSupplierPoQueryResult(v *AlibabaCeresSupplierPoQueryResult) {
	v.Values = v.Values[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaCeresSupplierPoQueryResult.Put(v)
}
