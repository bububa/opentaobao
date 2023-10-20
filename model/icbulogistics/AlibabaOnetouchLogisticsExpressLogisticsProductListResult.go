package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressLogisticsProductListResult 结构体
type AlibabaOnetouchLogisticsExpressLogisticsProductListResult struct {
	// 列表对象
	Values []LogisticsProductDto `json:"values,omitempty" xml:"values>logistics_product_dto,omitempty"`
	// 错误信息提示
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressLogisticsProductListResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressLogisticsProductListResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressLogisticsProductListResult() 从对象池中获取AlibabaOnetouchLogisticsExpressLogisticsProductListResult
func GetAlibabaOnetouchLogisticsExpressLogisticsProductListResult() *AlibabaOnetouchLogisticsExpressLogisticsProductListResult {
	return poolAlibabaOnetouchLogisticsExpressLogisticsProductListResult.Get().(*AlibabaOnetouchLogisticsExpressLogisticsProductListResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListResult 释放AlibabaOnetouchLogisticsExpressLogisticsProductListResult
func ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListResult(v *AlibabaOnetouchLogisticsExpressLogisticsProductListResult) {
	v.Values = v.Values[:0]
	v.ErrorMessage = ""
	v.ResultCode = 0
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressLogisticsProductListResult.Put(v)
}
