package promotion

import (
	"sync"
)

// ErrorMessage 结构体
type ErrorMessage struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 发送失败的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// asas
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

var poolErrorMessage = sync.Pool{
	New: func() any {
		return new(ErrorMessage)
	},
}

// GetErrorMessage() 从对象池中获取ErrorMessage
func GetErrorMessage() *ErrorMessage {
	return poolErrorMessage.Get().(*ErrorMessage)
}

// ReleaseErrorMessage 释放ErrorMessage
func ReleaseErrorMessage(v *ErrorMessage) {
	v.BuyerNick = ""
	v.Reason = ""
	v.Ouid = ""
	v.OpenUid = ""
	poolErrorMessage.Put(v)
}
