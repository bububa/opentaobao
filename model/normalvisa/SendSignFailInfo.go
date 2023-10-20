package normalvisa

import (
	"sync"
)

// SendSignFailInfo 结构体
type SendSignFailInfo struct {
	// 申请人id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolSendSignFailInfo = sync.Pool{
	New: func() any {
		return new(SendSignFailInfo)
	},
}

// GetSendSignFailInfo() 从对象池中获取SendSignFailInfo
func GetSendSignFailInfo() *SendSignFailInfo {
	return poolSendSignFailInfo.Get().(*SendSignFailInfo)
}

// ReleaseSendSignFailInfo 释放SendSignFailInfo
func ReleaseSendSignFailInfo(v *SendSignFailInfo) {
	v.ApplyId = ""
	v.ErrorMsg = ""
	poolSendSignFailInfo.Put(v)
}
