package ieagency

import (
	"sync"
)

// IeWriteBackPnrVo 结构体
type IeWriteBackPnrVo struct {
	// 预定订单pnr信息
	BookPnrVos []IeBookPnrVo `json:"book_pnr_vos,omitempty" xml:"book_pnr_vos>ie_book_pnr_vo,omitempty"`
	// 预定订单ID
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
}

var poolIeWriteBackPnrVo = sync.Pool{
	New: func() any {
		return new(IeWriteBackPnrVo)
	},
}

// GetIeWriteBackPnrVo() 从对象池中获取IeWriteBackPnrVo
func GetIeWriteBackPnrVo() *IeWriteBackPnrVo {
	return poolIeWriteBackPnrVo.Get().(*IeWriteBackPnrVo)
}

// ReleaseIeWriteBackPnrVo 释放IeWriteBackPnrVo
func ReleaseIeWriteBackPnrVo(v *IeWriteBackPnrVo) {
	v.BookPnrVos = v.BookPnrVos[:0]
	v.BookOrderId = 0
	poolIeWriteBackPnrVo.Put(v)
}
