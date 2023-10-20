package traveltrade

import (
	"sync"
)

// MemoUpdate 结构体
type MemoUpdate struct {
	// 关闭订单时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolMemoUpdate = sync.Pool{
	New: func() any {
		return new(MemoUpdate)
	},
}

// GetMemoUpdate() 从对象池中获取MemoUpdate
func GetMemoUpdate() *MemoUpdate {
	return poolMemoUpdate.Get().(*MemoUpdate)
}

// ReleaseMemoUpdate 释放MemoUpdate
func ReleaseMemoUpdate(v *MemoUpdate) {
	v.Modified = ""
	v.Tid = 0
	poolMemoUpdate.Put(v)
}
