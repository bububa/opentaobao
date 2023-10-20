package alitripmerchant

import (
	"sync"
)

// ReplaceRpCode 结构体
type ReplaceRpCode struct {
	// 代替后商品
	ReplaceRpCode string `json:"replace_rp_code,omitempty" xml:"replace_rp_code,omitempty"`
	// 原价商品
	OrigRpCode string `json:"orig_rp_code,omitempty" xml:"orig_rp_code,omitempty"`
}

var poolReplaceRpCode = sync.Pool{
	New: func() any {
		return new(ReplaceRpCode)
	},
}

// GetReplaceRpCode() 从对象池中获取ReplaceRpCode
func GetReplaceRpCode() *ReplaceRpCode {
	return poolReplaceRpCode.Get().(*ReplaceRpCode)
}

// ReleaseReplaceRpCode 释放ReplaceRpCode
func ReleaseReplaceRpCode(v *ReplaceRpCode) {
	v.ReplaceRpCode = ""
	v.OrigRpCode = ""
	poolReplaceRpCode.Put(v)
}
