package train

import (
	"sync"
)

// FreeChildrenTicketDetailRq 结构体
type FreeChildrenTicketDetailRq struct {
	// 申请单号
	ApplyNo int64 `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
}

var poolFreeChildrenTicketDetailRq = sync.Pool{
	New: func() any {
		return new(FreeChildrenTicketDetailRq)
	},
}

// GetFreeChildrenTicketDetailRq() 从对象池中获取FreeChildrenTicketDetailRq
func GetFreeChildrenTicketDetailRq() *FreeChildrenTicketDetailRq {
	return poolFreeChildrenTicketDetailRq.Get().(*FreeChildrenTicketDetailRq)
}

// ReleaseFreeChildrenTicketDetailRq 释放FreeChildrenTicketDetailRq
func ReleaseFreeChildrenTicketDetailRq(v *FreeChildrenTicketDetailRq) {
	v.ApplyNo = 0
	poolFreeChildrenTicketDetailRq.Put(v)
}
