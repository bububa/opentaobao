package baoxian

import (
	"sync"
)

// AliSceneResult 结构体
type AliSceneResult struct {
	// 系统自动生成
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 系统自动生成
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

var poolAliSceneResult = sync.Pool{
	New: func() any {
		return new(AliSceneResult)
	},
}

// GetAliSceneResult() 从对象池中获取AliSceneResult
func GetAliSceneResult() *AliSceneResult {
	return poolAliSceneResult.Get().(*AliSceneResult)
}

// ReleaseAliSceneResult 释放AliSceneResult
func ReleaseAliSceneResult(v *AliSceneResult) {
	v.BizErrorCode = ""
	v.BizErrorMsg = ""
	v.BizSuccess = false
	poolAliSceneResult.Put(v)
}
