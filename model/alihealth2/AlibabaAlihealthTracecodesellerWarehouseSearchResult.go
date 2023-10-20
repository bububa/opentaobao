package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTracecodesellerWarehouseSearchResult 结构体
type AlibabaAlihealthTracecodesellerWarehouseSearchResult struct {
	// detail
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolAlibabaAlihealthTracecodesellerWarehouseSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerWarehouseSearchResult)
	},
}

// GetAlibabaAlihealthTracecodesellerWarehouseSearchResult() 从对象池中获取AlibabaAlihealthTracecodesellerWarehouseSearchResult
func GetAlibabaAlihealthTracecodesellerWarehouseSearchResult() *AlibabaAlihealthTracecodesellerWarehouseSearchResult {
	return poolAlibabaAlihealthTracecodesellerWarehouseSearchResult.Get().(*AlibabaAlihealthTracecodesellerWarehouseSearchResult)
}

// ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchResult 释放AlibabaAlihealthTracecodesellerWarehouseSearchResult
func ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchResult(v *AlibabaAlihealthTracecodesellerWarehouseSearchResult) {
	v.Detail = ""
	v.Name = ""
	v.Code = ""
	v.Id = 0
	poolAlibabaAlihealthTracecodesellerWarehouseSearchResult.Put(v)
}
