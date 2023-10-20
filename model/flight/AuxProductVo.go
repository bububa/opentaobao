package flight

import (
	"sync"
)

// AuxProductVo 结构体
type AuxProductVo struct {
	// 外部outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 辅营产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 结构化的行李说明
	BaggageVo *BaggageVo `json:"baggage_vo,omitempty" xml:"baggage_vo,omitempty"`
	// 基准价
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 柜台价
	CounterPrice int64 `json:"counter_price,omitempty" xml:"counter_price,omitempty"`
	// 销售价格
	OnlinePrice int64 `json:"online_price,omitempty" xml:"online_price,omitempty"`
	// deprecated, 不建议使用
	ProductType int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 销售类型 1:一次销售，2:二次销售，3:不限
	SaleType int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 选座规格信息
	SeatVo *SeatVo `json:"seat_vo,omitempty" xml:"seat_vo,omitempty"`
	// 产品类型 4:付费行李，6:值机，7:餐食, 8:值机及选座
	UnityProductType int64 `json:"unity_product_type,omitempty" xml:"unity_product_type,omitempty"`
}

var poolAuxProductVo = sync.Pool{
	New: func() any {
		return new(AuxProductVo)
	},
}

// GetAuxProductVo() 从对象池中获取AuxProductVo
func GetAuxProductVo() *AuxProductVo {
	return poolAuxProductVo.Get().(*AuxProductVo)
}

// ReleaseAuxProductVo 释放AuxProductVo
func ReleaseAuxProductVo(v *AuxProductVo) {
	v.OuterId = ""
	v.ProductName = ""
	v.BaggageVo = nil
	v.BasePrice = 0
	v.CounterPrice = 0
	v.OnlinePrice = 0
	v.ProductType = 0
	v.SaleType = 0
	v.SeatVo = nil
	v.UnityProductType = 0
	poolAuxProductVo.Put(v)
}
