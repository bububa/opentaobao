package openmall

import (
	"sync"
)

// TaobaoOpenmallItemGetResultDo 结构体
type TaobaoOpenmallItemGetResultDo struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 商品
	Item *TopItemVo `json:"item,omitempty" xml:"item,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOpenmallItemGetResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemGetResultDo)
	},
}

// GetTaobaoOpenmallItemGetResultDo() 从对象池中获取TaobaoOpenmallItemGetResultDo
func GetTaobaoOpenmallItemGetResultDo() *TaobaoOpenmallItemGetResultDo {
	return poolTaobaoOpenmallItemGetResultDo.Get().(*TaobaoOpenmallItemGetResultDo)
}

// ReleaseTaobaoOpenmallItemGetResultDo 释放TaobaoOpenmallItemGetResultDo
func ReleaseTaobaoOpenmallItemGetResultDo(v *TaobaoOpenmallItemGetResultDo) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Item = nil
	v.Success = false
	poolTaobaoOpenmallItemGetResultDo.Put(v)
}
