package tvupadmin

import (
	"sync"
)

// DicControlApkDo 结构体
type DicControlApkDo struct {
}

var poolDicControlApkDo = sync.Pool{
	New: func() any {
		return new(DicControlApkDo)
	},
}

// GetDicControlApkDo() 从对象池中获取DicControlApkDo
func GetDicControlApkDo() *DicControlApkDo {
	return poolDicControlApkDo.Get().(*DicControlApkDo)
}

// ReleaseDicControlApkDo 释放DicControlApkDo
func ReleaseDicControlApkDo(v *DicControlApkDo) {
	poolDicControlApkDo.Put(v)
}
