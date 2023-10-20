package ieagency

import (
	"sync"
)

// ReceiveRefundTicketRs 结构体
type ReceiveRefundTicketRs struct {
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolReceiveRefundTicketRs = sync.Pool{
	New: func() any {
		return new(ReceiveRefundTicketRs)
	},
}

// GetReceiveRefundTicketRs() 从对象池中获取ReceiveRefundTicketRs
func GetReceiveRefundTicketRs() *ReceiveRefundTicketRs {
	return poolReceiveRefundTicketRs.Get().(*ReceiveRefundTicketRs)
}

// ReleaseReceiveRefundTicketRs 释放ReceiveRefundTicketRs
func ReleaseReceiveRefundTicketRs(v *ReceiveRefundTicketRs) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolReceiveRefundTicketRs.Put(v)
}
