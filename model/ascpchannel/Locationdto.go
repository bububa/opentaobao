package ascpchannel

import (
	"sync"
)

// Locationdto 结构体
type Locationdto struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}

var poolLocationdto = sync.Pool{
	New: func() any {
		return new(Locationdto)
	},
}

// GetLocationdto() 从对象池中获取Locationdto
func GetLocationdto() *Locationdto {
	return poolLocationdto.Get().(*Locationdto)
}

// ReleaseLocationdto 释放Locationdto
func ReleaseLocationdto(v *Locationdto) {
	v.StoreCode = ""
	poolLocationdto.Put(v)
}
