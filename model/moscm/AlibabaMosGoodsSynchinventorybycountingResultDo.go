package moscm

import (
	"sync"
)

// AlibabaMosGoodsSynchinventorybycountingResultDo 结构体
type AlibabaMosGoodsSynchinventorybycountingResultDo struct {
	// 返回数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaMosGoodsSynchinventorybycountingResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsSynchinventorybycountingResultDo)
	},
}

// GetAlibabaMosGoodsSynchinventorybycountingResultDo() 从对象池中获取AlibabaMosGoodsSynchinventorybycountingResultDo
func GetAlibabaMosGoodsSynchinventorybycountingResultDo() *AlibabaMosGoodsSynchinventorybycountingResultDo {
	return poolAlibabaMosGoodsSynchinventorybycountingResultDo.Get().(*AlibabaMosGoodsSynchinventorybycountingResultDo)
}

// ReleaseAlibabaMosGoodsSynchinventorybycountingResultDo 释放AlibabaMosGoodsSynchinventorybycountingResultDo
func ReleaseAlibabaMosGoodsSynchinventorybycountingResultDo(v *AlibabaMosGoodsSynchinventorybycountingResultDo) {
	v.Data = ""
	poolAlibabaMosGoodsSynchinventorybycountingResultDo.Put(v)
}
