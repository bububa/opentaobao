package alicom

import (
	"sync"
)

// SequenceCallRule 结构体
type SequenceCallRule struct {
	// 无振铃无法接通时，顺振到下一个号码，例如空号、停机、关机等。该参数不传时默认为1。1：生效
	CallNoRinging string `json:"call_no_ringing,omitempty" xml:"call_no_ringing,omitempty"`
	// 忙呼转的情况，指被叫本身设置了【呼叫转移号码】（181信令），出现呼叫转移时顺振到下一个号码。该参数不传时默认为1。1：生效
	CallForwarded string `json:"call_forwarded,omitempty" xml:"call_forwarded,omitempty"`
	// 忙拒绝情况，指【被叫忙】以及【被叫拒接】时，顺振到下一个号码。该参数不传时默认为1。1：生效
	BusyReject string `json:"busy_reject,omitempty" xml:"busy_reject,omitempty"`
	// 振铃无响应超时顺振,传入数字n，表示振铃后n秒后顺振到下一个号码。该参数不传时默认为30（也即是30秒顺振下一个号码），单位秒
	CallTimeout string `json:"call_timeout,omitempty" xml:"call_timeout,omitempty"`
}

var poolSequenceCallRule = sync.Pool{
	New: func() any {
		return new(SequenceCallRule)
	},
}

// GetSequenceCallRule() 从对象池中获取SequenceCallRule
func GetSequenceCallRule() *SequenceCallRule {
	return poolSequenceCallRule.Get().(*SequenceCallRule)
}

// ReleaseSequenceCallRule 释放SequenceCallRule
func ReleaseSequenceCallRule(v *SequenceCallRule) {
	v.CallNoRinging = ""
	v.CallForwarded = ""
	v.BusyReject = ""
	v.CallTimeout = ""
	poolSequenceCallRule.Put(v)
}
