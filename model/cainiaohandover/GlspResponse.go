package cainiaohandover

import (
	"sync"
)

// GlspResponse 结构体
type GlspResponse struct {
	// 错误信息
	ErrorInfo *ErrorInfo `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 请求结果
	Result *SolutionServiceResQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 查询是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolGlspResponse = sync.Pool{
	New: func() any {
		return new(GlspResponse)
	},
}

// GetGlspResponse() 从对象池中获取GlspResponse
func GetGlspResponse() *GlspResponse {
	return poolGlspResponse.Get().(*GlspResponse)
}

// ReleaseGlspResponse 释放GlspResponse
func ReleaseGlspResponse(v *GlspResponse) {
	v.ErrorInfo = nil
	v.Result = nil
	v.IsSuccess = false
	poolGlspResponse.Put(v)
}
