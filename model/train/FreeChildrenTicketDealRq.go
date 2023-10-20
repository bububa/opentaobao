package train

import (
	"sync"
)

// FreeChildrenTicketDealRq 结构体
type FreeChildrenTicketDealRq struct {
	// 失败文案
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 唯一id
	ApplyNo int64 `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
	// 失败code
	FailCode int64 `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolFreeChildrenTicketDealRq = sync.Pool{
	New: func() any {
		return new(FreeChildrenTicketDealRq)
	},
}

// GetFreeChildrenTicketDealRq() 从对象池中获取FreeChildrenTicketDealRq
func GetFreeChildrenTicketDealRq() *FreeChildrenTicketDealRq {
	return poolFreeChildrenTicketDealRq.Get().(*FreeChildrenTicketDealRq)
}

// ReleaseFreeChildrenTicketDealRq 释放FreeChildrenTicketDealRq
func ReleaseFreeChildrenTicketDealRq(v *FreeChildrenTicketDealRq) {
	v.FailMsg = ""
	v.ApplyNo = 0
	v.FailCode = 0
	v.IsSuccess = false
	poolFreeChildrenTicketDealRq.Put(v)
}
