package ascpchannel

import (
	"sync"
)

// Ownerdto 结构体
type Ownerdto struct {
	// 物流货主ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolOwnerdto = sync.Pool{
	New: func() any {
		return new(Ownerdto)
	},
}

// GetOwnerdto() 从对象池中获取Ownerdto
func GetOwnerdto() *Ownerdto {
	return poolOwnerdto.Get().(*Ownerdto)
}

// ReleaseOwnerdto 释放Ownerdto
func ReleaseOwnerdto(v *Ownerdto) {
	v.UserId = 0
	v.SupplierId = 0
	poolOwnerdto.Put(v)
}
