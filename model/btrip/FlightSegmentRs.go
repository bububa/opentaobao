package btrip

import (
	"sync"
)

// FlightSegmentRs 结构体
type FlightSegmentRs struct {
	// 航班到达机场三字码
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 到达城市三字码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 到达国家
	ArrCountry string `json:"arr_country,omitempty" xml:"arr_country,omitempty"`
	// 航班到达航站楼
	ArrTerm string `json:"arr_term,omitempty" xml:"arr_term,omitempty"`
	// 到达日期时间(yyyy-MM-dd HH:mm:ss)
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 航班起飞机场三字码
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 出发城市三字码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发国家
	DepCountry string `json:"dep_country,omitempty" xml:"dep_country,omitempty"`
	// 航班出发航站楼
	DepTerm string `json:"dep_term,omitempty" xml:"dep_term,omitempty"`
	// 出发日期时间(yyyy-MM-dd HH:mm:ss)
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 机型
	EquipType string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// 航班号+出发机场+达到机场+起飞时间（精确分）
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 市场方航司(如：KA)
	MarketingAirline string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// 市场方航班号(如：KA5809)
	MarketingFlightNo string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// 经停城市，多个值使用&#34;,&#34;分隔
	StopCity string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// 到达数字日期(yyyyMMdd)
	ArrDateInt int64 `json:"arr_date_int,omitempty" xml:"arr_date_int,omitempty"`
	// 出发数字日期(yyyyMMdd)
	DepDateInt int64 `json:"dep_date_int,omitempty" xml:"dep_date_int,omitempty"`
	// 时长(单位：分钟)
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 市场方数字航班号（如：5809）
	MarketingFlightNoInt int64 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// 餐食类型,0:无参食；1：有餐食；2：饮品；3：茶点；4：早餐；5：正餐；
	Meal int64 `json:"meal,omitempty" xml:"meal,omitempty"`
	// 里程
	Miles int64 `json:"miles,omitempty" xml:"miles,omitempty"`
	// 航段展示信息
	SegmentShowInfo *SegmentShowInfoRs `json:"segment_show_info,omitempty" xml:"segment_show_info,omitempty"`
	// 航段序号，从0开始
	SeqId int64 `json:"seq_id,omitempty" xml:"seq_id,omitempty"`
	// 经停次数
	StopQuantity int64 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
	// 时间
	TransferTime int64 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
	// 是否换机场
	TransferChangeAirport bool `json:"transfer_change_airport,omitempty" xml:"transfer_change_airport,omitempty"`
}

var poolFlightSegmentRs = sync.Pool{
	New: func() any {
		return new(FlightSegmentRs)
	},
}

// GetFlightSegmentRs() 从对象池中获取FlightSegmentRs
func GetFlightSegmentRs() *FlightSegmentRs {
	return poolFlightSegmentRs.Get().(*FlightSegmentRs)
}

// ReleaseFlightSegmentRs 释放FlightSegmentRs
func ReleaseFlightSegmentRs(v *FlightSegmentRs) {
	v.ArrAirport = ""
	v.ArrCity = ""
	v.ArrCountry = ""
	v.ArrTerm = ""
	v.ArrTime = ""
	v.DepAirport = ""
	v.DepCity = ""
	v.DepCountry = ""
	v.DepTerm = ""
	v.DepTime = ""
	v.EquipType = ""
	v.Id = ""
	v.MarketingAirline = ""
	v.MarketingFlightNo = ""
	v.StopCity = ""
	v.ArrDateInt = 0
	v.DepDateInt = 0
	v.Duration = 0
	v.MarketingFlightNoInt = 0
	v.Meal = 0
	v.Miles = 0
	v.SegmentShowInfo = nil
	v.SeqId = 0
	v.StopQuantity = 0
	v.TransferTime = 0
	v.TransferChangeAirport = false
	poolFlightSegmentRs.Put(v)
}
