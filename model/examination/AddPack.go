package examination

import (
	"sync"
)

// AddPack 结构体
type AddPack struct {
	// 加项包id
	IsvPackId string `json:"isv_pack_id,omitempty" xml:"isv_pack_id,omitempty"`
	// 版本号，isv需要进行校验版本号是否过期，判断加项包信息是否已更改，健康未同步
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolAddPack = sync.Pool{
	New: func() any {
		return new(AddPack)
	},
}

// GetAddPack() 从对象池中获取AddPack
func GetAddPack() *AddPack {
	return poolAddPack.Get().(*AddPack)
}

// ReleaseAddPack 释放AddPack
func ReleaseAddPack(v *AddPack) {
	v.IsvPackId = ""
	v.Version = 0
	poolAddPack.Put(v)
}
