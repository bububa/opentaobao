package ascpchannel

import (
	"sync"
)

// Packagequeryrequest 结构体
type Packagequeryrequest struct {
	// 仓库编码
	WlbOrderCode string `json:"wlb_order_code,omitempty" xml:"wlb_order_code,omitempty"`
}

var poolPackagequeryrequest = sync.Pool{
	New: func() any {
		return new(Packagequeryrequest)
	},
}

// GetPackagequeryrequest() 从对象池中获取Packagequeryrequest
func GetPackagequeryrequest() *Packagequeryrequest {
	return poolPackagequeryrequest.Get().(*Packagequeryrequest)
}

// ReleasePackagequeryrequest 释放Packagequeryrequest
func ReleasePackagequeryrequest(v *Packagequeryrequest) {
	v.WlbOrderCode = ""
	poolPackagequeryrequest.Put(v)
}
