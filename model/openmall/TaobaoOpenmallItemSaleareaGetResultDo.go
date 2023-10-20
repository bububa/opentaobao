package openmall

import (
	"sync"
)

// TaobaoOpenmallItemSaleareaGetResultDo 结构体
type TaobaoOpenmallItemSaleareaGetResultDo struct {
	// 可售区域结果
	SaleAreaList []TopSaleAreaVo `json:"sale_area_list,omitempty" xml:"sale_area_list>top_sale_area_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOpenmallItemSaleareaGetResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemSaleareaGetResultDo)
	},
}

// GetTaobaoOpenmallItemSaleareaGetResultDo() 从对象池中获取TaobaoOpenmallItemSaleareaGetResultDo
func GetTaobaoOpenmallItemSaleareaGetResultDo() *TaobaoOpenmallItemSaleareaGetResultDo {
	return poolTaobaoOpenmallItemSaleareaGetResultDo.Get().(*TaobaoOpenmallItemSaleareaGetResultDo)
}

// ReleaseTaobaoOpenmallItemSaleareaGetResultDo 释放TaobaoOpenmallItemSaleareaGetResultDo
func ReleaseTaobaoOpenmallItemSaleareaGetResultDo(v *TaobaoOpenmallItemSaleareaGetResultDo) {
	v.SaleAreaList = v.SaleAreaList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoOpenmallItemSaleareaGetResultDo.Put(v)
}
