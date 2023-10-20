package openmall

import (
	"sync"
)

// TopParseAddressVo 结构体
type TopParseAddressVo struct {
	// 地址解析结构
	Entries []TopParseAddressEntryVo `json:"entries,omitempty" xml:"entries>top_parse_address_entry_vo,omitempty"`
}

var poolTopParseAddressVo = sync.Pool{
	New: func() any {
		return new(TopParseAddressVo)
	},
}

// GetTopParseAddressVo() 从对象池中获取TopParseAddressVo
func GetTopParseAddressVo() *TopParseAddressVo {
	return poolTopParseAddressVo.Get().(*TopParseAddressVo)
}

// ReleaseTopParseAddressVo 释放TopParseAddressVo
func ReleaseTopParseAddressVo(v *TopParseAddressVo) {
	v.Entries = v.Entries[:0]
	poolTopParseAddressVo.Put(v)
}
