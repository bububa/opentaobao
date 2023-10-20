package simba

import (
	"sync"
)

// GuidancePrice 结构体
type GuidancePrice struct {
	// 相对当前价格点击提升比例
	ClickUpRate string `json:"click_up_rate,omitempty" xml:"click_up_rate,omitempty"`
	// 相对当前价格展现提升比例
	ImpressionUpRate string `json:"impression_up_rate,omitempty" xml:"impression_up_rate,omitempty"`
	// 价格类型，0代表正常排名建议，1代表被过滤排名建议， 		 * 2代表相近价格建议，3代表转化出价建议
	PriceFlag string `json:"price_flag,omitempty" xml:"price_flag,omitempty"`
	// 建议出价（分）
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// flag＝0，1时为目标排名，flag＝2时无意义,flag=3时代表对应的建议
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 建议价格能够拿到的点击量
	NewClick string `json:"new_click,omitempty" xml:"new_click,omitempty"`
	// 建议价格能够拿到的展现量
	NewImpression string `json:"new_impression,omitempty" xml:"new_impression,omitempty"`
}

var poolGuidancePrice = sync.Pool{
	New: func() any {
		return new(GuidancePrice)
	},
}

// GetGuidancePrice() 从对象池中获取GuidancePrice
func GetGuidancePrice() *GuidancePrice {
	return poolGuidancePrice.Get().(*GuidancePrice)
}

// ReleaseGuidancePrice 释放GuidancePrice
func ReleaseGuidancePrice(v *GuidancePrice) {
	v.ClickUpRate = ""
	v.ImpressionUpRate = ""
	v.PriceFlag = ""
	v.Price = ""
	v.Flag = ""
	v.NewClick = ""
	v.NewImpression = ""
	poolGuidancePrice.Put(v)
}
