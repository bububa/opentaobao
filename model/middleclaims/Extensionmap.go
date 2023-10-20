package middleclaims

import (
	"sync"
)

// Extensionmap 结构体
type Extensionmap struct {
	// 预留扩展值1
	ExtensionMapFir string `json:"extension_map_fir,omitempty" xml:"extension_map_fir,omitempty"`
	// 预留扩展值2
	ExtensionMapSec string `json:"extension_map_sec,omitempty" xml:"extension_map_sec,omitempty"`
	// 预留扩展值3
	ExtensionMapThi string `json:"extension_map_thi,omitempty" xml:"extension_map_thi,omitempty"`
	// 预留扩展Map1
	ExtensionMap string `json:"extension_map,omitempty" xml:"extension_map,omitempty"`
}

var poolExtensionmap = sync.Pool{
	New: func() any {
		return new(Extensionmap)
	},
}

// GetExtensionmap() 从对象池中获取Extensionmap
func GetExtensionmap() *Extensionmap {
	return poolExtensionmap.Get().(*Extensionmap)
}

// ReleaseExtensionmap 释放Extensionmap
func ReleaseExtensionmap(v *Extensionmap) {
	v.ExtensionMapFir = ""
	v.ExtensionMapSec = ""
	v.ExtensionMapThi = ""
	v.ExtensionMap = ""
	poolExtensionmap.Put(v)
}
