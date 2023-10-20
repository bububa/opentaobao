package traveltrade

import (
	"sync"
)

// MemoCreate 结构体
type MemoCreate struct {
	// 备注添加时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolMemoCreate = sync.Pool{
	New: func() any {
		return new(MemoCreate)
	},
}

// GetMemoCreate() 从对象池中获取MemoCreate
func GetMemoCreate() *MemoCreate {
	return poolMemoCreate.Get().(*MemoCreate)
}

// ReleaseMemoCreate 释放MemoCreate
func ReleaseMemoCreate(v *MemoCreate) {
	v.Created = ""
	v.Tid = 0
	poolMemoCreate.Put(v)
}
