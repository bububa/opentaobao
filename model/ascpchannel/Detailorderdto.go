package ascpchannel

import (
	"sync"
)

// Detailorderdto 结构体
type Detailorderdto struct {
	// 业务发生时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 实际操作子单id
	OperationDetailOrderId string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
}

var poolDetailorderdto = sync.Pool{
	New: func() any {
		return new(Detailorderdto)
	},
}

// GetDetailorderdto() 从对象池中获取Detailorderdto
func GetDetailorderdto() *Detailorderdto {
	return poolDetailorderdto.Get().(*Detailorderdto)
}

// ReleaseDetailorderdto 释放Detailorderdto
func ReleaseDetailorderdto(v *Detailorderdto) {
	v.BizDate = ""
	v.OperationDetailOrderId = ""
	poolDetailorderdto.Put(v)
}
