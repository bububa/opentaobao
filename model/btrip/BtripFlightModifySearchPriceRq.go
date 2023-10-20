package btrip

import (
	"sync"
)

// BtripFlightModifySearchPriceRq 结构体
type BtripFlightModifySearchPriceRq struct {
	// 乘客列表
	TravelerInfoList []TravelerInfo `json:"traveler_info_list,omitempty" xml:"traveler_info_list>traveler_info,omitempty"`
	// 到达城市三字码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 出发城市三字码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 外部订单
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 供应商标识
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 0:非自愿 1:自愿
	IsVoluntary int64 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
}

var poolBtripFlightModifySearchPriceRq = sync.Pool{
	New: func() any {
		return new(BtripFlightModifySearchPriceRq)
	},
}

// GetBtripFlightModifySearchPriceRq() 从对象池中获取BtripFlightModifySearchPriceRq
func GetBtripFlightModifySearchPriceRq() *BtripFlightModifySearchPriceRq {
	return poolBtripFlightModifySearchPriceRq.Get().(*BtripFlightModifySearchPriceRq)
}

// ReleaseBtripFlightModifySearchPriceRq 释放BtripFlightModifySearchPriceRq
func ReleaseBtripFlightModifySearchPriceRq(v *BtripFlightModifySearchPriceRq) {
	v.TravelerInfoList = v.TravelerInfoList[:0]
	v.ArrCity = ""
	v.CabinClass = ""
	v.DepCity = ""
	v.DepDate = ""
	v.FlightNo = ""
	v.DisOrderId = ""
	v.SubChannel = ""
	v.SupplierCode = ""
	v.IsVoluntary = 0
	poolBtripFlightModifySearchPriceRq.Put(v)
}
