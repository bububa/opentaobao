package smartstore

import (
	"sync"
)

// TmallPopupstoreItemDiscountPriceResultDto 结构体
type TmallPopupstoreItemDiscountPriceResultDto struct {
	// 错误码code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 实际结果
	ResultList string `json:"result_list,omitempty" xml:"result_list,omitempty"`
	// 数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolTmallPopupstoreItemDiscountPriceResultDto = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreItemDiscountPriceResultDto)
	},
}

// GetTmallPopupstoreItemDiscountPriceResultDto() 从对象池中获取TmallPopupstoreItemDiscountPriceResultDto
func GetTmallPopupstoreItemDiscountPriceResultDto() *TmallPopupstoreItemDiscountPriceResultDto {
	return poolTmallPopupstoreItemDiscountPriceResultDto.Get().(*TmallPopupstoreItemDiscountPriceResultDto)
}

// ReleaseTmallPopupstoreItemDiscountPriceResultDto 释放TmallPopupstoreItemDiscountPriceResultDto
func ReleaseTmallPopupstoreItemDiscountPriceResultDto(v *TmallPopupstoreItemDiscountPriceResultDto) {
	v.Code = ""
	v.Msg = ""
	v.ResultList = ""
	v.Total = 0
	poolTmallPopupstoreItemDiscountPriceResultDto.Put(v)
}
