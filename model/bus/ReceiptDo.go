package bus

import (
	"sync"
)

// ReceiptDo 结构体
type ReceiptDo struct {
	// 开票时间
	ReceiptDateTime string `json:"receipt_date_time,omitempty" xml:"receipt_date_time,omitempty"`
	// 发票流水号。成功时必填
	ReceiptNumber string `json:"receipt_number,omitempty" xml:"receipt_number,omitempty"`
	// 发票链接.status为1时必填
	ReceiptUrl string `json:"receipt_url,omitempty" xml:"receipt_url,omitempty"`
	// 失败原因.status为0时必填
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 失败码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 发票状态1成功0失败-1异常
	ReceiptStatus int64 `json:"receipt_status,omitempty" xml:"receipt_status,omitempty"`
}

var poolReceiptDo = sync.Pool{
	New: func() any {
		return new(ReceiptDo)
	},
}

// GetReceiptDo() 从对象池中获取ReceiptDo
func GetReceiptDo() *ReceiptDo {
	return poolReceiptDo.Get().(*ReceiptDo)
}

// ReleaseReceiptDo 释放ReceiptDo
func ReleaseReceiptDo(v *ReceiptDo) {
	v.ReceiptDateTime = ""
	v.ReceiptNumber = ""
	v.ReceiptUrl = ""
	v.FailReason = ""
	v.FailCode = ""
	v.AgentId = 0
	v.ReceiptStatus = 0
	poolReceiptDo.Put(v)
}
