package flight

import (
	"sync"
)

// Segments 结构体
type Segments struct {
	// 舱等  (0:头等舱 1:商务舱 2:经济舱 3:超值经济舱 4:标准经济舱 5:超级经济舱)
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 起飞时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 到达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 出发机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 航班实际承运航司
	OperatingAirline string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// 共享航班
	OperatingFlightNo string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// 航段序号
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 航程序号
	OdIndex int64 `json:"od_index,omitempty" xml:"od_index,omitempty"`
	// 公布票面价
	SegPatPrice int64 `json:"seg_pat_price,omitempty" xml:"seg_pat_price,omitempty"`
}

var poolSegments = sync.Pool{
	New: func() any {
		return new(Segments)
	},
}

// GetSegments() 从对象池中获取Segments
func GetSegments() *Segments {
	return poolSegments.Get().(*Segments)
}

// ReleaseSegments 释放Segments
func ReleaseSegments(v *Segments) {
	v.CabinClass = ""
	v.FlightNo = ""
	v.DepTime = ""
	v.ArrCity = ""
	v.DepCity = ""
	v.Cabin = ""
	v.ArrAirport = ""
	v.DepAirport = ""
	v.ArrTime = ""
	v.OperatingAirline = ""
	v.OperatingFlightNo = ""
	v.SegmentIndex = 0
	v.OdIndex = 0
	v.SegPatPrice = 0
	poolSegments.Put(v)
}
