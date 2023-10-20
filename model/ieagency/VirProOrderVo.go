package ieagency

import (
	"sync"
)

// VirProOrderVo 结构体
type VirProOrderVo struct {
	// bookTime
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolVirProOrderVo = sync.Pool{
	New: func() any {
		return new(VirProOrderVo)
	},
}

// GetVirProOrderVo() 从对象池中获取VirProOrderVo
func GetVirProOrderVo() *VirProOrderVo {
	return poolVirProOrderVo.Get().(*VirProOrderVo)
}

// ReleaseVirProOrderVo 释放VirProOrderVo
func ReleaseVirProOrderVo(v *VirProOrderVo) {
	v.BookTime = ""
	v.OrderId = 0
	poolVirProOrderVo.Put(v)
}
