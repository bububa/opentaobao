package nlife

import (
	"sync"
)

// DiscountMemo 结构体
type DiscountMemo struct {
	// 文案编号
	MemoId string `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
	// 文案描述
	MemoDesc string `json:"memo_desc,omitempty" xml:"memo_desc,omitempty"`
}

var poolDiscountMemo = sync.Pool{
	New: func() any {
		return new(DiscountMemo)
	},
}

// GetDiscountMemo() 从对象池中获取DiscountMemo
func GetDiscountMemo() *DiscountMemo {
	return poolDiscountMemo.Get().(*DiscountMemo)
}

// ReleaseDiscountMemo 释放DiscountMemo
func ReleaseDiscountMemo(v *DiscountMemo) {
	v.MemoId = ""
	v.MemoDesc = ""
	poolDiscountMemo.Put(v)
}
