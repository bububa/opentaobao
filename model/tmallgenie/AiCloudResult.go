package tmallgenie

import (
	"sync"
)

// AiCloudResult 结构体
type AiCloudResult struct {
	// 返回item列表
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// 语音地址
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAiCloudResult = sync.Pool{
	New: func() any {
		return new(AiCloudResult)
	},
}

// GetAiCloudResult() 从对象池中获取AiCloudResult
func GetAiCloudResult() *AiCloudResult {
	return poolAiCloudResult.Get().(*AiCloudResult)
}

// ReleaseAiCloudResult 释放AiCloudResult
func ReleaseAiCloudResult(v *AiCloudResult) {
	v.Models = v.Models[:0]
	v.Model = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.IsSuccess = false
	v.Success = false
	poolAiCloudResult.Put(v)
}
