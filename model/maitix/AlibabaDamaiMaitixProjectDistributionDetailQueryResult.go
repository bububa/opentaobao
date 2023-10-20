package maitix

import (
	"sync"
)

// AlibabaDamaiMaitixProjectDistributionDetailQueryResult 结构体
type AlibabaDamaiMaitixProjectDistributionDetailQueryResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *OpenProjectDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMaitixProjectDistributionDetailQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixProjectDistributionDetailQueryResult)
	},
}

// GetAlibabaDamaiMaitixProjectDistributionDetailQueryResult() 从对象池中获取AlibabaDamaiMaitixProjectDistributionDetailQueryResult
func GetAlibabaDamaiMaitixProjectDistributionDetailQueryResult() *AlibabaDamaiMaitixProjectDistributionDetailQueryResult {
	return poolAlibabaDamaiMaitixProjectDistributionDetailQueryResult.Get().(*AlibabaDamaiMaitixProjectDistributionDetailQueryResult)
}

// ReleaseAlibabaDamaiMaitixProjectDistributionDetailQueryResult 释放AlibabaDamaiMaitixProjectDistributionDetailQueryResult
func ReleaseAlibabaDamaiMaitixProjectDistributionDetailQueryResult(v *AlibabaDamaiMaitixProjectDistributionDetailQueryResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaDamaiMaitixProjectDistributionDetailQueryResult.Put(v)
}
