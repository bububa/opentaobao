package eleenterprisecoupon

import (
	"sync"
)

// CartCouponDto 结构体
type CartCouponDto struct {
	// 券描述信息
	Descriptions []string `json:"descriptions,omitempty" xml:"descriptions>string,omitempty"`
	// 券不可用原因列表，该字段仅在购物车场景有效
	UnavailableReasons []string `json:"unavailable_reasons,omitempty" xml:"unavailable_reasons>string,omitempty"`
	// 红包的SN或者券的ID
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// 券名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠类型为立减时表示立减金额, 优惠类型为特价券时，表示特价金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 使用门槛
	Threshold string `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// 券剩余张数
	StockCountLeft string `json:"stock_count_left,omitempty" xml:"stock_count_left,omitempty"`
	// 类型：红包=1，券=2
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 券的张数，默认是1，ET时表示使用的张数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 券的展示状态  1：券可用可选中  2：券可用但是不可选中  3：券不可用，不可选中
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 券可使用张数
	StockCountAvailable int64 `json:"stock_count_available,omitempty" xml:"stock_count_available,omitempty"`
}

var poolCartCouponDto = sync.Pool{
	New: func() any {
		return new(CartCouponDto)
	},
}

// GetCartCouponDto() 从对象池中获取CartCouponDto
func GetCartCouponDto() *CartCouponDto {
	return poolCartCouponDto.Get().(*CartCouponDto)
}

// ReleaseCartCouponDto 释放CartCouponDto
func ReleaseCartCouponDto(v *CartCouponDto) {
	v.Descriptions = v.Descriptions[:0]
	v.UnavailableReasons = v.UnavailableReasons[:0]
	v.Sn = ""
	v.Name = ""
	v.Amount = ""
	v.Threshold = ""
	v.StockCountLeft = ""
	v.PromotionType = 0
	v.Quantity = 0
	v.Status = 0
	v.StockCountAvailable = 0
	poolCartCouponDto.Put(v)
}
