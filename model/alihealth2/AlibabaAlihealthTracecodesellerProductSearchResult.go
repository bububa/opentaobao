package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTracecodesellerProductSearchResult 结构体
type AlibabaAlihealthTracecodesellerProductSearchResult struct {
	// 商品备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 商品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// productTracInfoId
	ProductTracInfoId int64 `json:"product_trac_info_id,omitempty" xml:"product_trac_info_id,omitempty"`
}

var poolAlibabaAlihealthTracecodesellerProductSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerProductSearchResult)
	},
}

// GetAlibabaAlihealthTracecodesellerProductSearchResult() 从对象池中获取AlibabaAlihealthTracecodesellerProductSearchResult
func GetAlibabaAlihealthTracecodesellerProductSearchResult() *AlibabaAlihealthTracecodesellerProductSearchResult {
	return poolAlibabaAlihealthTracecodesellerProductSearchResult.Get().(*AlibabaAlihealthTracecodesellerProductSearchResult)
}

// ReleaseAlibabaAlihealthTracecodesellerProductSearchResult 释放AlibabaAlihealthTracecodesellerProductSearchResult
func ReleaseAlibabaAlihealthTracecodesellerProductSearchResult(v *AlibabaAlihealthTracecodesellerProductSearchResult) {
	v.Remark = ""
	v.ProductName = ""
	v.ProductTracInfoId = 0
	poolAlibabaAlihealthTracecodesellerProductSearchResult.Put(v)
}
