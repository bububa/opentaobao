package tbtrade

import (
	"sync"
)

// CallBackResultDto 结构体
type CallBackResultDto struct {
	// 0：成功，-1： 缺少必填字段 ，-2： 参数解析错误 ，-3： 内部处理错误
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 请求传入的tecAgent
	TecAgent string `json:"tec_agent,omitempty" xml:"tec_agent,omitempty"`
	// 请求传入的appAccessToken
	AppAccessToken string `json:"app_access_token,omitempty" xml:"app_access_token,omitempty"`
}

var poolCallBackResultDto = sync.Pool{
	New: func() any {
		return new(CallBackResultDto)
	},
}

// GetCallBackResultDto() 从对象池中获取CallBackResultDto
func GetCallBackResultDto() *CallBackResultDto {
	return poolCallBackResultDto.Get().(*CallBackResultDto)
}

// ReleaseCallBackResultDto 释放CallBackResultDto
func ReleaseCallBackResultDto(v *CallBackResultDto) {
	v.Result = ""
	v.TecAgent = ""
	v.AppAccessToken = ""
	poolCallBackResultDto.Put(v)
}
