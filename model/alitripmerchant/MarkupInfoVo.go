package alitripmerchant

import (
	"sync"
)

// MarkupInfoVo 结构体
type MarkupInfoVo struct {
	// 每日加价信息
	DailyMarkupPriceList []DailyMarkupPrice `json:"daily_markup_price_list,omitempty" xml:"daily_markup_price_list>daily_markup_price,omitempty"`
	// 加价描述
	MarkupDesc string `json:"markup_desc,omitempty" xml:"markup_desc,omitempty"`
	// 加价金额
	MarkupPrice int64 `json:"markup_price,omitempty" xml:"markup_price,omitempty"`
	// 加价金额总和
	TotalMarkupPrice int64 `json:"total_markup_price,omitempty" xml:"total_markup_price,omitempty"`
}

var poolMarkupInfoVo = sync.Pool{
	New: func() any {
		return new(MarkupInfoVo)
	},
}

// GetMarkupInfoVo() 从对象池中获取MarkupInfoVo
func GetMarkupInfoVo() *MarkupInfoVo {
	return poolMarkupInfoVo.Get().(*MarkupInfoVo)
}

// ReleaseMarkupInfoVo 释放MarkupInfoVo
func ReleaseMarkupInfoVo(v *MarkupInfoVo) {
	v.DailyMarkupPriceList = v.DailyMarkupPriceList[:0]
	v.MarkupDesc = ""
	v.MarkupPrice = 0
	v.TotalMarkupPrice = 0
	poolMarkupInfoVo.Put(v)
}
