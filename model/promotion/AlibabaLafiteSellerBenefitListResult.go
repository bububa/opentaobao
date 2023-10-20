package promotion

import (
	"sync"
)

// AlibabaLafiteSellerBenefitListResult 结构体
type AlibabaLafiteSellerBenefitListResult struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回查询结果
	Data *Page `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLafiteSellerBenefitListResult = sync.Pool{
	New: func() any {
		return new(AlibabaLafiteSellerBenefitListResult)
	},
}

// GetAlibabaLafiteSellerBenefitListResult() 从对象池中获取AlibabaLafiteSellerBenefitListResult
func GetAlibabaLafiteSellerBenefitListResult() *AlibabaLafiteSellerBenefitListResult {
	return poolAlibabaLafiteSellerBenefitListResult.Get().(*AlibabaLafiteSellerBenefitListResult)
}

// ReleaseAlibabaLafiteSellerBenefitListResult 释放AlibabaLafiteSellerBenefitListResult
func ReleaseAlibabaLafiteSellerBenefitListResult(v *AlibabaLafiteSellerBenefitListResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLafiteSellerBenefitListResult.Put(v)
}
