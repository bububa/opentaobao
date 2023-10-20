package nrt

import (
	"sync"
)

// TmallNrtItemMainSynchronizeResultDo 结构体
type TmallNrtItemMainSynchronizeResultDo struct {
	// 返回值
	Data *NrtItemSyncResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolTmallNrtItemMainSynchronizeResultDo = sync.Pool{
	New: func() any {
		return new(TmallNrtItemMainSynchronizeResultDo)
	},
}

// GetTmallNrtItemMainSynchronizeResultDo() 从对象池中获取TmallNrtItemMainSynchronizeResultDo
func GetTmallNrtItemMainSynchronizeResultDo() *TmallNrtItemMainSynchronizeResultDo {
	return poolTmallNrtItemMainSynchronizeResultDo.Get().(*TmallNrtItemMainSynchronizeResultDo)
}

// ReleaseTmallNrtItemMainSynchronizeResultDo 释放TmallNrtItemMainSynchronizeResultDo
func ReleaseTmallNrtItemMainSynchronizeResultDo(v *TmallNrtItemMainSynchronizeResultDo) {
	v.Data = nil
	v.Succ = false
	poolTmallNrtItemMainSynchronizeResultDo.Put(v)
}
