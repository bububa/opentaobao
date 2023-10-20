package qianniu

import (
	"sync"
)

// QnCopilotResultDo 结构体
type QnCopilotResultDo struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 审核结果
	Result *PicAuditResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolQnCopilotResultDo = sync.Pool{
	New: func() any {
		return new(QnCopilotResultDo)
	},
}

// GetQnCopilotResultDo() 从对象池中获取QnCopilotResultDo
func GetQnCopilotResultDo() *QnCopilotResultDo {
	return poolQnCopilotResultDo.Get().(*QnCopilotResultDo)
}

// ReleaseQnCopilotResultDo 释放QnCopilotResultDo
func ReleaseQnCopilotResultDo(v *QnCopilotResultDo) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	poolQnCopilotResultDo.Put(v)
}
