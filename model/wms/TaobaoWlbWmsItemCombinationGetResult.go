package wms

import (
	"sync"
)

// TaobaoWlbWmsItemCombinationGetResult 结构体
type TaobaoWlbWmsItemCombinationGetResult struct {
	// 子货品列表
	SubItemList []SubItemList `json:"sub_item_list,omitempty" xml:"sub_item_list>sub_item_list,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

var poolTaobaoWlbWmsItemCombinationGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsItemCombinationGetResult)
	},
}

// GetTaobaoWlbWmsItemCombinationGetResult() 从对象池中获取TaobaoWlbWmsItemCombinationGetResult
func GetTaobaoWlbWmsItemCombinationGetResult() *TaobaoWlbWmsItemCombinationGetResult {
	return poolTaobaoWlbWmsItemCombinationGetResult.Get().(*TaobaoWlbWmsItemCombinationGetResult)
}

// ReleaseTaobaoWlbWmsItemCombinationGetResult 释放TaobaoWlbWmsItemCombinationGetResult
func ReleaseTaobaoWlbWmsItemCombinationGetResult(v *TaobaoWlbWmsItemCombinationGetResult) {
	v.SubItemList = v.SubItemList[:0]
	v.WlErrorCode = ""
	v.WlErrorMsg = ""
	v.WlSuccess = false
	poolTaobaoWlbWmsItemCombinationGetResult.Put(v)
}
