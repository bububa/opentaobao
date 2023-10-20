package xhotelonlineorder

import (
	"sync"
)

// TopInvoiceDo 结构体
type TopInvoiceDo struct {
	// 发票提供方:0未知 1酒店前台，2卖家开具
	ProviderType int64 `json:"provider_type,omitempty" xml:"provider_type,omitempty"`
}

var poolTopInvoiceDo = sync.Pool{
	New: func() any {
		return new(TopInvoiceDo)
	},
}

// GetTopInvoiceDo() 从对象池中获取TopInvoiceDo
func GetTopInvoiceDo() *TopInvoiceDo {
	return poolTopInvoiceDo.Get().(*TopInvoiceDo)
}

// ReleaseTopInvoiceDo 释放TopInvoiceDo
func ReleaseTopInvoiceDo(v *TopInvoiceDo) {
	v.ProviderType = 0
	poolTopInvoiceDo.Put(v)
}
