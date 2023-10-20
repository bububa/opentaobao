package btrip

import (
	"sync"
)

// BtripTrainOpenTicketDo 结构体
type BtripTrainOpenTicketDo struct {
	// 车次号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 车次类型
	TrainType string `json:"train_type,omitempty" xml:"train_type,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 运行时长
	RunTime string `json:"run_time,omitempty" xml:"run_time,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发车站
	DepStation string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// 到达车站
	ArrStation string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// 坐席
	Seat string `json:"seat,omitempty" xml:"seat,omitempty"`
	// 乘客名字
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 票价（分）
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// tmc收取的服务费（分）
	ServiceFee int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 原票价（分）
	OriginTicketPrice int64 `json:"origin_ticket_price,omitempty" xml:"origin_ticket_price,omitempty"`
	// 是否退改
	IsChanged bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
}

var poolBtripTrainOpenTicketDo = sync.Pool{
	New: func() any {
		return new(BtripTrainOpenTicketDo)
	},
}

// GetBtripTrainOpenTicketDo() 从对象池中获取BtripTrainOpenTicketDo
func GetBtripTrainOpenTicketDo() *BtripTrainOpenTicketDo {
	return poolBtripTrainOpenTicketDo.Get().(*BtripTrainOpenTicketDo)
}

// ReleaseBtripTrainOpenTicketDo 释放BtripTrainOpenTicketDo
func ReleaseBtripTrainOpenTicketDo(v *BtripTrainOpenTicketDo) {
	v.TrainNo = ""
	v.TrainType = ""
	v.DepTime = ""
	v.ArrTime = ""
	v.RunTime = ""
	v.DepCity = ""
	v.ArrCity = ""
	v.DepStation = ""
	v.ArrStation = ""
	v.Seat = ""
	v.PassengerName = ""
	v.TicketPrice = 0
	v.ServiceFee = 0
	v.OriginTicketPrice = 0
	v.IsChanged = false
	poolBtripTrainOpenTicketDo.Put(v)
}
