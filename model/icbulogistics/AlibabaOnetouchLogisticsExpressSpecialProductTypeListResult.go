package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult 结构体
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult struct {
	// 列表对象
	Values []SpecialProductTypeDto `json:"values,omitempty" xml:"values>special_product_type_dto,omitempty"`
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult() 从对象池中获取AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult
func GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult {
	return poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult.Get().(*AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult 释放AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult
func ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult(v *AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult) {
	v.Values = v.Values[:0]
	v.ErrorMessage = ""
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListResult.Put(v)
}
