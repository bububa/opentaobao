package btrip

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// BtripFlightCreateOrderRq 结构体
type BtripFlightCreateOrderRq struct {
	// 航段列表
	FlightSegmentList []BookFlightSegmentDto `json:"flight_segment_list,omitempty" xml:"flight_segment_list>book_flight_segment_dto,omitempty"`
	// 出行人
	TravelerInfoList []TravelerInfo `json:"traveler_info_list,omitempty" xml:"traveler_info_list>traveler_info,omitempty"`
	// 出发城市
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 预订人名字
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 预订人员工号
	BuyerUniqueKey string `json:"buyer_unique_key,omitempty" xml:"buyer_unique_key,omitempty"`
	// 出发城市
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 扩展字段
	OrderAttr string `json:"order_attr,omitempty" xml:"order_attr,omitempty"`
	// 邮寄地址
	ReceiptAddress string `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	// 发票抬头
	ReceiptTitle string `json:"receipt_title,omitempty" xml:"receipt_title,omitempty"`
	// 搜索下单参数
	OrderParams string `json:"order_params,omitempty" xml:"order_params,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 外部订单ID
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 出发机场编码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 到达机场编码
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 商品id
	OtaItemId string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// 是否自动支付
	AutoPay *model.File `json:"auto_pay,omitempty" xml:"auto_pay,omitempty"`
	// 联系人信息
	ContactInfo *ContactInfoDto `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// 2000
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 发票邮寄类型
	ReceiptTarget int64 `json:"receipt_target,omitempty" xml:"receipt_target,omitempty"`
	// 航程类型
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 是否异步创单
	AsyncCreateOrderMode bool `json:"async_create_order_mode,omitempty" xml:"async_create_order_mode,omitempty"`
}

var poolBtripFlightCreateOrderRq = sync.Pool{
	New: func() any {
		return new(BtripFlightCreateOrderRq)
	},
}

// GetBtripFlightCreateOrderRq() 从对象池中获取BtripFlightCreateOrderRq
func GetBtripFlightCreateOrderRq() *BtripFlightCreateOrderRq {
	return poolBtripFlightCreateOrderRq.Get().(*BtripFlightCreateOrderRq)
}

// ReleaseBtripFlightCreateOrderRq 释放BtripFlightCreateOrderRq
func ReleaseBtripFlightCreateOrderRq(v *BtripFlightCreateOrderRq) {
	v.FlightSegmentList = v.FlightSegmentList[:0]
	v.TravelerInfoList = v.TravelerInfoList[:0]
	v.ArrCityCode = ""
	v.BuyerName = ""
	v.BuyerUniqueKey = ""
	v.DepCityCode = ""
	v.DepDate = ""
	v.OrderAttr = ""
	v.ReceiptAddress = ""
	v.ReceiptTitle = ""
	v.OrderParams = ""
	v.SubChannel = ""
	v.DisOrderId = ""
	v.DepAirportCode = ""
	v.ArrAirportCode = ""
	v.OtaItemId = ""
	v.AutoPay = nil
	v.ContactInfo = nil
	v.Price = 0
	v.ReceiptTarget = 0
	v.TripType = 0
	v.AsyncCreateOrderMode = false
	poolBtripFlightCreateOrderRq.Put(v)
}
