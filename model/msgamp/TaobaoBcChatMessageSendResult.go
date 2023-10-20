package msgamp

import (
	"sync"
)

// TaobaoBcChatMessageSendResult 结构体
type TaobaoBcChatMessageSendResult struct {
	// SERVICE_ERROR
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// SERVICE_ERROR
	MsgErrMessage string `json:"msg_err_message,omitempty" xml:"msg_err_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTaobaoBcChatMessageSendResult = sync.Pool{
	New: func() any {
		return new(TaobaoBcChatMessageSendResult)
	},
}

// GetTaobaoBcChatMessageSendResult() 从对象池中获取TaobaoBcChatMessageSendResult
func GetTaobaoBcChatMessageSendResult() *TaobaoBcChatMessageSendResult {
	return poolTaobaoBcChatMessageSendResult.Get().(*TaobaoBcChatMessageSendResult)
}

// ReleaseTaobaoBcChatMessageSendResult 释放TaobaoBcChatMessageSendResult
func ReleaseTaobaoBcChatMessageSendResult(v *TaobaoBcChatMessageSendResult) {
	v.MsgCode = ""
	v.Model = ""
	v.MsgErrMessage = ""
	v.IsSuccess = false
	poolTaobaoBcChatMessageSendResult.Put(v)
}
