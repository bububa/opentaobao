package opentrade

import (
	"sync"
)

// ItemUnBindResult 结构体
type ItemUnBindResult struct {
	// 解绑异常时的错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否解绑成功
	BindOk bool `json:"bind_ok,omitempty" xml:"bind_ok,omitempty"`
}

var poolItemUnBindResult = sync.Pool{
	New: func() any {
		return new(ItemUnBindResult)
	},
}

// GetItemUnBindResult() 从对象池中获取ItemUnBindResult
func GetItemUnBindResult() *ItemUnBindResult {
	return poolItemUnBindResult.Get().(*ItemUnBindResult)
}

// ReleaseItemUnBindResult 释放ItemUnBindResult
func ReleaseItemUnBindResult(v *ItemUnBindResult) {
	v.ErrorMessage = ""
	v.ItemId = 0
	v.BindOk = false
	poolItemUnBindResult.Put(v)
}
