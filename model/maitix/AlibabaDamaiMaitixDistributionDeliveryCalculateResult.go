package maitix

import (
	"sync"
)

// AlibabaDamaiMaitixDistributionDeliveryCalculateResult 结构体
type AlibabaDamaiMaitixDistributionDeliveryCalculateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务对象
	Model *OpenApiPostFeeDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMaitixDistributionDeliveryCalculateResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionDeliveryCalculateResult)
	},
}

// GetAlibabaDamaiMaitixDistributionDeliveryCalculateResult() 从对象池中获取AlibabaDamaiMaitixDistributionDeliveryCalculateResult
func GetAlibabaDamaiMaitixDistributionDeliveryCalculateResult() *AlibabaDamaiMaitixDistributionDeliveryCalculateResult {
	return poolAlibabaDamaiMaitixDistributionDeliveryCalculateResult.Get().(*AlibabaDamaiMaitixDistributionDeliveryCalculateResult)
}

// ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateResult 释放AlibabaDamaiMaitixDistributionDeliveryCalculateResult
func ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateResult(v *AlibabaDamaiMaitixDistributionDeliveryCalculateResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = nil
	v.Success = false
	poolAlibabaDamaiMaitixDistributionDeliveryCalculateResult.Put(v)
}
