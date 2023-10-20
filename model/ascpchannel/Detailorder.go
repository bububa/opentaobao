package ascpchannel

import (
	"sync"
)

// Detailorder 结构体
type Detailorder struct {
	// 实际操作子单id(例如:ICP子单,,UDP子单)
	OperationDetailOrderId string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
}

var poolDetailorder = sync.Pool{
	New: func() any {
		return new(Detailorder)
	},
}

// GetDetailorder() 从对象池中获取Detailorder
func GetDetailorder() *Detailorder {
	return poolDetailorder.Get().(*Detailorder)
}

// ReleaseDetailorder 释放Detailorder
func ReleaseDetailorder(v *Detailorder) {
	v.OperationDetailOrderId = ""
	poolDetailorder.Put(v)
}
