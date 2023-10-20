package ascpchannel

import (
	"sync"
)

// Logisticsacceptinforeslist 结构体
type Logisticsacceptinforeslist struct {
	// 菜鸟订单编码
	PresalesOrderId string `json:"presales_order_id,omitempty" xml:"presales_order_id,omitempty"`
}

var poolLogisticsacceptinforeslist = sync.Pool{
	New: func() any {
		return new(Logisticsacceptinforeslist)
	},
}

// GetLogisticsacceptinforeslist() 从对象池中获取Logisticsacceptinforeslist
func GetLogisticsacceptinforeslist() *Logisticsacceptinforeslist {
	return poolLogisticsacceptinforeslist.Get().(*Logisticsacceptinforeslist)
}

// ReleaseLogisticsacceptinforeslist 释放Logisticsacceptinforeslist
func ReleaseLogisticsacceptinforeslist(v *Logisticsacceptinforeslist) {
	v.PresalesOrderId = ""
	poolLogisticsacceptinforeslist.Put(v)
}
