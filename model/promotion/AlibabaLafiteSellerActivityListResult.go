package promotion

import (
	"sync"
)

// AlibabaLafiteSellerActivityListResult 结构体
type AlibabaLafiteSellerActivityListResult struct {
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口返回结果
	PageData *Page `json:"page_data,omitempty" xml:"page_data,omitempty"`
}

var poolAlibabaLafiteSellerActivityListResult = sync.Pool{
	New: func() any {
		return new(AlibabaLafiteSellerActivityListResult)
	},
}

// GetAlibabaLafiteSellerActivityListResult() 从对象池中获取AlibabaLafiteSellerActivityListResult
func GetAlibabaLafiteSellerActivityListResult() *AlibabaLafiteSellerActivityListResult {
	return poolAlibabaLafiteSellerActivityListResult.Get().(*AlibabaLafiteSellerActivityListResult)
}

// ReleaseAlibabaLafiteSellerActivityListResult 释放AlibabaLafiteSellerActivityListResult
func ReleaseAlibabaLafiteSellerActivityListResult(v *AlibabaLafiteSellerActivityListResult) {
	v.Msg = ""
	v.Code = ""
	v.PageData = nil
	poolAlibabaLafiteSellerActivityListResult.Put(v)
}
