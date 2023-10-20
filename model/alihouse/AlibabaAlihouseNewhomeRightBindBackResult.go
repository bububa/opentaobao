package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeRightBindBackResult 结构体
type AlibabaAlihouseNewhomeRightBindBackResult struct {
	// 1
	Data []AstoreSceneRespInfoDto `json:"data,omitempty" xml:"data>astore_scene_resp_info_dto,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeRightBindBackResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRightBindBackResult)
	},
}

// GetAlibabaAlihouseNewhomeRightBindBackResult() 从对象池中获取AlibabaAlihouseNewhomeRightBindBackResult
func GetAlibabaAlihouseNewhomeRightBindBackResult() *AlibabaAlihouseNewhomeRightBindBackResult {
	return poolAlibabaAlihouseNewhomeRightBindBackResult.Get().(*AlibabaAlihouseNewhomeRightBindBackResult)
}

// ReleaseAlibabaAlihouseNewhomeRightBindBackResult 释放AlibabaAlihouseNewhomeRightBindBackResult
func ReleaseAlibabaAlihouseNewhomeRightBindBackResult(v *AlibabaAlihouseNewhomeRightBindBackResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeRightBindBackResult.Put(v)
}
