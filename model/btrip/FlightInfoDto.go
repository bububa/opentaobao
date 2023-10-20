package btrip

import (
	"sync"
)

// FlightInfoDto 结构体
type FlightInfoDto struct {
	// 最低舱位价格列表
	LowestCabinPrice []ModifyPrice `json:"lowest_cabin_price,omitempty" xml:"lowest_cabin_price>modify_price,omitempty"`
	// 舱位列表
	CabinList []CabinInfo `json:"cabin_list,omitempty" xml:"cabin_list>cabin_info,omitempty"`
	// 最低舱位价格列表
	LowestCabinPriceList []ModifyPrice `json:"lowest_cabin_price_list,omitempty" xml:"lowest_cabin_price_list>modify_price,omitempty"`
	// 机票退改规则等规则列表
	FlightRuleList []FlightRule `json:"flight_rule_list,omitempty" xml:"flight_rule_list>flight_rule,omitempty"`
	// 多舱位价格
	CabinInfoList []MultiCabinClassInfo `json:"cabin_info_list,omitempty" xml:"cabin_info_list>multi_cabin_class_info,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 共享航班号（共享航班时有值）
	CarrierNo string `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	// 共享航班（共享航班时有值）
	CarrierAirline string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// 最低舱数量
	LowestCabinNum string `json:"lowest_cabin_num,omitempty" xml:"lowest_cabin_num,omitempty"`
	// 最低舱位
	LowestCabin string `json:"lowest_cabin,omitempty" xml:"lowest_cabin,omitempty"`
	// 最低舱等
	LowestCabinClass string `json:"lowest_cabin_class,omitempty" xml:"lowest_cabin_class,omitempty"`
	// 改签航班出发日期
	ModifyFlightDepDate string `json:"modify_flight_dep_date,omitempty" xml:"modify_flight_dep_date,omitempty"`
	// 改签航班出发时间
	ModifyFlightDepTime string `json:"modify_flight_dep_time,omitempty" xml:"modify_flight_dep_time,omitempty"`
	// 改签航班到达时间
	ModifyFlightArrTime string `json:"modify_flight_arr_time,omitempty" xml:"modify_flight_arr_time,omitempty"`
	// 会话Id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 出发城市三字码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 经停到达时间
	StopArrTime string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// 经停城市
	StopCity string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// 经停出发时间
	StopDepTime string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// 总价
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 规则描述
	ClassRule string `json:"class_rule,omitempty" xml:"class_rule,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 优惠金额
	PromotionPrice string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 剩余座位数
	RemainedSeatCount string `json:"remained_seat_count,omitempty" xml:"remained_seat_count,omitempty"`
	// 加密参数
	SecretParams string `json:"secret_params,omitempty" xml:"secret_params,omitempty"`
	// 航段值
	SegmentNumber string `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	// 到达日期
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 餐食
	MealDesc string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// 产品类型描述
	ProductTypeDesc string `json:"product_type_desc,omitempty" xml:"product_type_desc,omitempty"`
	// 机型
	FlightSize string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	// 机型代码
	FlightType string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// 最低价舱位描述
	LowestCabinDesc string `json:"lowest_cabin_desc,omitempty" xml:"lowest_cabin_desc,omitempty"`
	// 承运航司信息
	CarrierAirLine string `json:"carrier_air_line,omitempty" xml:"carrier_air_line,omitempty"`
	// 商品id
	OtaItemId string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// 航司信息
	AirlineInfo *AirlineInfo `json:"airline_info,omitempty" xml:"airline_info,omitempty"`
	// 出发机场信息
	DepAirportInfo *AirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty"`
	// 到达机场信息
	ArrAirportInfo *AirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty"`
	// 机建费
	BuildPrice int64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// 折扣
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 燃油费用
	OilPrice int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// 票价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 发票类型
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 行程类型
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 中转信息
	TransferInfo *TransferInfo `json:"transfer_info,omitempty" xml:"transfer_info,omitempty"`
	// 票面价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 舱位基准价
	BasicCabinPrice int64 `json:"basic_cabin_price,omitempty" xml:"basic_cabin_price,omitempty"`
	// 是否共享
	IsShare bool `json:"is_share,omitempty" xml:"is_share,omitempty"`
	// 是否经停
	IsStop bool `json:"is_stop,omitempty" xml:"is_stop,omitempty"`
	// 是否中转
	IsTransfer bool `json:"is_transfer,omitempty" xml:"is_transfer,omitempty"`
	// 是否协议价
	IsProtocol bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
}

var poolFlightInfoDto = sync.Pool{
	New: func() any {
		return new(FlightInfoDto)
	},
}

// GetFlightInfoDto() 从对象池中获取FlightInfoDto
func GetFlightInfoDto() *FlightInfoDto {
	return poolFlightInfoDto.Get().(*FlightInfoDto)
}

// ReleaseFlightInfoDto 释放FlightInfoDto
func ReleaseFlightInfoDto(v *FlightInfoDto) {
	v.LowestCabinPrice = v.LowestCabinPrice[:0]
	v.CabinList = v.CabinList[:0]
	v.LowestCabinPriceList = v.LowestCabinPriceList[:0]
	v.FlightRuleList = v.FlightRuleList[:0]
	v.CabinInfoList = v.CabinInfoList[:0]
	v.FlightNo = ""
	v.CarrierNo = ""
	v.CarrierAirline = ""
	v.LowestCabinNum = ""
	v.LowestCabin = ""
	v.LowestCabinClass = ""
	v.ModifyFlightDepDate = ""
	v.ModifyFlightDepTime = ""
	v.ModifyFlightArrTime = ""
	v.SessionId = ""
	v.Cabin = ""
	v.CabinClass = ""
	v.DepCityCode = ""
	v.DepDate = ""
	v.StopArrTime = ""
	v.StopCity = ""
	v.StopDepTime = ""
	v.TotalPrice = ""
	v.ClassRule = ""
	v.Memo = ""
	v.PromotionPrice = ""
	v.RemainedSeatCount = ""
	v.SecretParams = ""
	v.SegmentNumber = ""
	v.ArrDate = ""
	v.MealDesc = ""
	v.ProductTypeDesc = ""
	v.FlightSize = ""
	v.FlightType = ""
	v.LowestCabinDesc = ""
	v.CarrierAirLine = ""
	v.OtaItemId = ""
	v.AirlineInfo = nil
	v.DepAirportInfo = nil
	v.ArrAirportInfo = nil
	v.BuildPrice = 0
	v.Discount = 0
	v.OilPrice = 0
	v.TicketPrice = 0
	v.InvoiceType = 0
	v.TripType = 0
	v.TransferInfo = nil
	v.Price = 0
	v.BasicCabinPrice = 0
	v.IsShare = false
	v.IsStop = false
	v.IsTransfer = false
	v.IsProtocol = false
	poolFlightInfoDto.Put(v)
}
