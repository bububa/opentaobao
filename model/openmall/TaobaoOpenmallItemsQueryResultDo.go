package openmall

import (
	"sync"
)

// TaobaoOpenmallItemsQueryResultDo 结构体
type TaobaoOpenmallItemsQueryResultDo struct {
	// 商品列表
	ItemList []TopItemVo `json:"item_list,omitempty" xml:"item_list>top_item_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOpenmallItemsQueryResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemsQueryResultDo)
	},
}

// GetTaobaoOpenmallItemsQueryResultDo() 从对象池中获取TaobaoOpenmallItemsQueryResultDo
func GetTaobaoOpenmallItemsQueryResultDo() *TaobaoOpenmallItemsQueryResultDo {
	return poolTaobaoOpenmallItemsQueryResultDo.Get().(*TaobaoOpenmallItemsQueryResultDo)
}

// ReleaseTaobaoOpenmallItemsQueryResultDo 释放TaobaoOpenmallItemsQueryResultDo
func ReleaseTaobaoOpenmallItemsQueryResultDo(v *TaobaoOpenmallItemsQueryResultDo) {
	v.ItemList = v.ItemList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoOpenmallItemsQueryResultDo.Put(v)
}
