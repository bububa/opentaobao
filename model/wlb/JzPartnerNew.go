package wlb

import (
	"sync"
)

// JzPartnerNew 结构体
type JzPartnerNew struct {
	// 服务商编码
	TpCode string `json:"tp_code,omitempty" xml:"tp_code,omitempty"`
	// 服务商名称
	TpName string `json:"tp_name,omitempty" xml:"tp_name,omitempty"`
	// 服务类型
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 是否是虚拟服务商（家装-商家自有物流）
	IsVirtualTp bool `json:"is_virtual_tp,omitempty" xml:"is_virtual_tp,omitempty"`
}

var poolJzPartnerNew = sync.Pool{
	New: func() any {
		return new(JzPartnerNew)
	},
}

// GetJzPartnerNew() 从对象池中获取JzPartnerNew
func GetJzPartnerNew() *JzPartnerNew {
	return poolJzPartnerNew.Get().(*JzPartnerNew)
}

// ReleaseJzPartnerNew 释放JzPartnerNew
func ReleaseJzPartnerNew(v *JzPartnerNew) {
	v.TpCode = ""
	v.TpName = ""
	v.ServiceType = 0
	v.IsVirtualTp = false
	poolJzPartnerNew.Put(v)
}
