package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressAddressStreetListResult 结构体
type AlibabaOnetouchLogisticsExpressAddressStreetListResult struct {
	// 列表对象
	Values []RegionEntity `json:"values,omitempty" xml:"values>region_entity,omitempty"`
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressAddressStreetListResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressStreetListResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressStreetListResult() 从对象池中获取AlibabaOnetouchLogisticsExpressAddressStreetListResult
func GetAlibabaOnetouchLogisticsExpressAddressStreetListResult() *AlibabaOnetouchLogisticsExpressAddressStreetListResult {
	return poolAlibabaOnetouchLogisticsExpressAddressStreetListResult.Get().(*AlibabaOnetouchLogisticsExpressAddressStreetListResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListResult 释放AlibabaOnetouchLogisticsExpressAddressStreetListResult
func ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListResult(v *AlibabaOnetouchLogisticsExpressAddressStreetListResult) {
	v.Values = v.Values[:0]
	v.ErrorMessage = ""
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressAddressStreetListResult.Put(v)
}
