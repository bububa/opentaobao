package alihealthmedical

import (
	"sync"
)

// OuterMsgContent 结构体
type OuterMsgContent struct {
	// 图片
	Pic []string `json:"pic,omitempty" xml:"pic>string,omitempty"`
	// 文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 音频
	Radio string `json:"radio,omitempty" xml:"radio,omitempty"`
	// 诊断
	Diagnose string `json:"diagnose,omitempty" xml:"diagnose,omitempty"`
	// 建议
	Advice string `json:"advice,omitempty" xml:"advice,omitempty"`
	// 音频时长
	RadioTime int64 `json:"radio_time,omitempty" xml:"radio_time,omitempty"`
	// 发送时间戳
	SendTime int64 `json:"send_time,omitempty" xml:"send_time,omitempty"`
}

var poolOuterMsgContent = sync.Pool{
	New: func() any {
		return new(OuterMsgContent)
	},
}

// GetOuterMsgContent() 从对象池中获取OuterMsgContent
func GetOuterMsgContent() *OuterMsgContent {
	return poolOuterMsgContent.Get().(*OuterMsgContent)
}

// ReleaseOuterMsgContent 释放OuterMsgContent
func ReleaseOuterMsgContent(v *OuterMsgContent) {
	v.Pic = v.Pic[:0]
	v.Text = ""
	v.Radio = ""
	v.Diagnose = ""
	v.Advice = ""
	v.RadioTime = 0
	v.SendTime = 0
	poolOuterMsgContent.Put(v)
}
