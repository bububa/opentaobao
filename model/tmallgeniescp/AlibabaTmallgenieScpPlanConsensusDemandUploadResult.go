package tmallgeniescp

import (
	"sync"
)

// AlibabaTmallgenieScpPlanConsensusDemandUploadResult 结构体
type AlibabaTmallgenieScpPlanConsensusDemandUploadResult struct {
	// 结果描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

var poolAlibabaTmallgenieScpPlanConsensusDemandUploadResult = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanConsensusDemandUploadResult)
	},
}

// GetAlibabaTmallgenieScpPlanConsensusDemandUploadResult() 从对象池中获取AlibabaTmallgenieScpPlanConsensusDemandUploadResult
func GetAlibabaTmallgenieScpPlanConsensusDemandUploadResult() *AlibabaTmallgenieScpPlanConsensusDemandUploadResult {
	return poolAlibabaTmallgenieScpPlanConsensusDemandUploadResult.Get().(*AlibabaTmallgenieScpPlanConsensusDemandUploadResult)
}

// ReleaseAlibabaTmallgenieScpPlanConsensusDemandUploadResult 释放AlibabaTmallgenieScpPlanConsensusDemandUploadResult
func ReleaseAlibabaTmallgenieScpPlanConsensusDemandUploadResult(v *AlibabaTmallgenieScpPlanConsensusDemandUploadResult) {
	v.Msg = ""
	v.Code = ""
	v.Suc = false
	poolAlibabaTmallgenieScpPlanConsensusDemandUploadResult.Put(v)
}
