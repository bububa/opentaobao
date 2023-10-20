package alicom

import (
	"sync"
)

// ControlRespDto 结构体
type ControlRespDto struct {
	// 接续控制信息:CONTINUE(接续),REJECT(拦截),IVR(收取用户键盘输入内容)
	ControlOperate string `json:"control_operate,omitempty" xml:"control_operate,omitempty"`
	// controlMsg
	ControlMsg string `json:"control_msg,omitempty" xml:"control_msg,omitempty"`
	// 对应到小号平台的能力类型:AXB、AXN、AXN_EXTENSION_REUSE(AXN分机复用)
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 主叫放音编码，多个文件用英文逗号分隔。
	CallNoPlayCode string `json:"call_no_play_code,omitempty" xml:"call_no_play_code,omitempty"`
	// 被叫放音编码，多个文件用英文逗号分隔。
	CalledNoPlayCode string `json:"called_no_play_code,omitempty" xml:"called_no_play_code,omitempty"`
	// 摘机后主叫侧的放音编码，多个文件用英文逗号分隔。
	CalledNoCallerPlayCode string `json:"called_no_caller_play_code,omitempty" xml:"called_no_caller_play_code,omitempty"`
	// 对应到阿里侧的绑定信息
	Subs *Subs `json:"subs,omitempty" xml:"subs,omitempty"`
	// 通话持续时长，可选，单位秒，如果出现则通话有效时长为此值，如果没有出现按现在默认处理
	CallDuration int64 `json:"call_duration,omitempty" xml:"call_duration,omitempty"`
	// 是否媒体资源降级,放弃录音放音功能；接入方无此相关功能，可忽略
	MediaDegrade bool `json:"media_degrade,omitempty" xml:"media_degrade,omitempty"`
}

var poolControlRespDto = sync.Pool{
	New: func() any {
		return new(ControlRespDto)
	},
}

// GetControlRespDto() 从对象池中获取ControlRespDto
func GetControlRespDto() *ControlRespDto {
	return poolControlRespDto.Get().(*ControlRespDto)
}

// ReleaseControlRespDto 释放ControlRespDto
func ReleaseControlRespDto(v *ControlRespDto) {
	v.ControlOperate = ""
	v.ControlMsg = ""
	v.ProductType = ""
	v.CallNoPlayCode = ""
	v.CalledNoPlayCode = ""
	v.CalledNoCallerPlayCode = ""
	v.Subs = nil
	v.CallDuration = 0
	v.MediaDegrade = false
	poolControlRespDto.Put(v)
}
