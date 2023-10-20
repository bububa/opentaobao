package ieagency

import (
	"sync"
)

// QueryRefundTicketsRs 结构体
type QueryRefundTicketsRs struct {
	// refundTicketList
	RefundTicketList []IeRefundTicketVo `json:"refund_ticket_list,omitempty" xml:"refund_ticket_list>ie_refund_ticket_vo,omitempty"`
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// pageCount
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolQueryRefundTicketsRs = sync.Pool{
	New: func() any {
		return new(QueryRefundTicketsRs)
	},
}

// GetQueryRefundTicketsRs() 从对象池中获取QueryRefundTicketsRs
func GetQueryRefundTicketsRs() *QueryRefundTicketsRs {
	return poolQueryRefundTicketsRs.Get().(*QueryRefundTicketsRs)
}

// ReleaseQueryRefundTicketsRs 释放QueryRefundTicketsRs
func ReleaseQueryRefundTicketsRs(v *QueryRefundTicketsRs) {
	v.RefundTicketList = v.RefundTicketList[:0]
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.PageCount = 0
	v.Success = false
	poolQueryRefundTicketsRs.Put(v)
}
