package simba

import (
	"sync"
)

// AdgroupOcpcVo 结构体
type AdgroupOcpcVo struct {
	// OCPC溢价比例
	OcpcRatio int64 `json:"ocpc_ratio,omitempty" xml:"ocpc_ratio,omitempty"`
	// OCPC是否开启，false:否，true:是
	EnableOcpc bool `json:"enable_ocpc,omitempty" xml:"enable_ocpc,omitempty"`
}

var poolAdgroupOcpcVo = sync.Pool{
	New: func() any {
		return new(AdgroupOcpcVo)
	},
}

// GetAdgroupOcpcVo() 从对象池中获取AdgroupOcpcVo
func GetAdgroupOcpcVo() *AdgroupOcpcVo {
	return poolAdgroupOcpcVo.Get().(*AdgroupOcpcVo)
}

// ReleaseAdgroupOcpcVo 释放AdgroupOcpcVo
func ReleaseAdgroupOcpcVo(v *AdgroupOcpcVo) {
	v.OcpcRatio = 0
	v.EnableOcpc = false
	poolAdgroupOcpcVo.Put(v)
}
