package travel

import (
	"sync"
)

// ProductSaleInfo 结构体
type ProductSaleInfo struct {
	// 建议零售价上限，单位：元
	PriceUpper string `json:"price_upper,omitempty" xml:"price_upper,omitempty"`
	// 建议零售价下限，单位：元
	PriceLower string `json:"price_lower,omitempty" xml:"price_lower,omitempty"`
	// 可选出发结束日期，格式：yyyy-MM-dd
	EndComboDate string `json:"end_combo_date,omitempty" xml:"end_combo_date,omitempty"`
	// 可选出发开始日期，格式：yyyy-MM-dd
	StartComboDate string `json:"start_combo_date,omitempty" xml:"start_combo_date,omitempty"`
	// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	ConfirmTime int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 产品线
	ProductLine int64 `json:"product_line,omitempty" xml:"product_line,omitempty"`
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	ConfirmType int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 商品售卖类型，0为日历商品，1为预约商品，2为非日历商品；默认为0
	SaleType int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 是否支持经销
	Distribute bool `json:"distribute,omitempty" xml:"distribute,omitempty"`
	// 是否支持代销
	Agent bool `json:"agent,omitempty" xml:"agent,omitempty"`
}

var poolProductSaleInfo = sync.Pool{
	New: func() any {
		return new(ProductSaleInfo)
	},
}

// GetProductSaleInfo() 从对象池中获取ProductSaleInfo
func GetProductSaleInfo() *ProductSaleInfo {
	return poolProductSaleInfo.Get().(*ProductSaleInfo)
}

// ReleaseProductSaleInfo 释放ProductSaleInfo
func ReleaseProductSaleInfo(v *ProductSaleInfo) {
	v.PriceUpper = ""
	v.PriceLower = ""
	v.EndComboDate = ""
	v.StartComboDate = ""
	v.ConfirmTime = 0
	v.Duration = 0
	v.ProductLine = 0
	v.ConfirmType = 0
	v.SaleType = 0
	v.Distribute = false
	v.Agent = false
	poolProductSaleInfo.Put(v)
}
