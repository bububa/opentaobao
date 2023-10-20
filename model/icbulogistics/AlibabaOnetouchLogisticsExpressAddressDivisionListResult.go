package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressAddressDivisionListResult 结构体
type AlibabaOnetouchLogisticsExpressAddressDivisionListResult struct {
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 列表对象
	Values *RegionEntity `json:"values,omitempty" xml:"values,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressAddressDivisionListResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressDivisionListResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressDivisionListResult() 从对象池中获取AlibabaOnetouchLogisticsExpressAddressDivisionListResult
func GetAlibabaOnetouchLogisticsExpressAddressDivisionListResult() *AlibabaOnetouchLogisticsExpressAddressDivisionListResult {
	return poolAlibabaOnetouchLogisticsExpressAddressDivisionListResult.Get().(*AlibabaOnetouchLogisticsExpressAddressDivisionListResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListResult 释放AlibabaOnetouchLogisticsExpressAddressDivisionListResult
func ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListResult(v *AlibabaOnetouchLogisticsExpressAddressDivisionListResult) {
	v.ErrorMessage = ""
	v.Values = nil
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressAddressDivisionListResult.Put(v)
}
