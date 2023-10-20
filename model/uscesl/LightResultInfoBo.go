package uscesl

import (
	"sync"
)

// LightResultInfoBo 结构体
type LightResultInfoBo struct {
	// 通知消息
	NotifyMessage string `json:"notify_message,omitempty" xml:"notify_message,omitempty"`
	// 成功数量
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
	// 失败数量
	FailCount int64 `json:"fail_count,omitempty" xml:"fail_count,omitempty"`
}

var poolLightResultInfoBo = sync.Pool{
	New: func() any {
		return new(LightResultInfoBo)
	},
}

// GetLightResultInfoBo() 从对象池中获取LightResultInfoBo
func GetLightResultInfoBo() *LightResultInfoBo {
	return poolLightResultInfoBo.Get().(*LightResultInfoBo)
}

// ReleaseLightResultInfoBo 释放LightResultInfoBo
func ReleaseLightResultInfoBo(v *LightResultInfoBo) {
	v.NotifyMessage = ""
	v.SuccessCount = 0
	v.FailCount = 0
	poolLightResultInfoBo.Put(v)
}
