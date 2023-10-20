package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeRightUnbindBackResult 结构体
type AlibabaAlihouseNewhomeRightUnbindBackResult struct {
	// 1
	Data []AstoreSceneRespInfoDto `json:"data,omitempty" xml:"data>astore_scene_resp_info_dto,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeRightUnbindBackResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRightUnbindBackResult)
	},
}

// GetAlibabaAlihouseNewhomeRightUnbindBackResult() 从对象池中获取AlibabaAlihouseNewhomeRightUnbindBackResult
func GetAlibabaAlihouseNewhomeRightUnbindBackResult() *AlibabaAlihouseNewhomeRightUnbindBackResult {
	return poolAlibabaAlihouseNewhomeRightUnbindBackResult.Get().(*AlibabaAlihouseNewhomeRightUnbindBackResult)
}

// ReleaseAlibabaAlihouseNewhomeRightUnbindBackResult 释放AlibabaAlihouseNewhomeRightUnbindBackResult
func ReleaseAlibabaAlihouseNewhomeRightUnbindBackResult(v *AlibabaAlihouseNewhomeRightUnbindBackResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeRightUnbindBackResult.Put(v)
}
