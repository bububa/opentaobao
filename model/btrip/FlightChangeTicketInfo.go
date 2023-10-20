package btrip

import (
	"sync"
)

// FlightChangeTicketInfo 结构体
type FlightChangeTicketInfo struct {
	// 改签票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 改签原票号
	OriginTicketNo string `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	// 改签航班号
	ChangeFlightNo string `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	// 改签舱位
	ChangeCabin string `json:"change_cabin,omitempty" xml:"change_cabin,omitempty"`
	// 改签舱等
	ChangeCabinLevel string `json:"change_cabin_level,omitempty" xml:"change_cabin_level,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 改签原因
	ChangeReason string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// 改签单id
	ChangeOrderId int64 `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// 改签类型：0自愿/1非自愿
	ChangeType int64 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// 机票改签费
	ChangeFee float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 机票升舱费
	UpgradeFee float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

var poolFlightChangeTicketInfo = sync.Pool{
	New: func() any {
		return new(FlightChangeTicketInfo)
	},
}

// GetFlightChangeTicketInfo() 从对象池中获取FlightChangeTicketInfo
func GetFlightChangeTicketInfo() *FlightChangeTicketInfo {
	return poolFlightChangeTicketInfo.Get().(*FlightChangeTicketInfo)
}

// ReleaseFlightChangeTicketInfo 释放FlightChangeTicketInfo
func ReleaseFlightChangeTicketInfo(v *FlightChangeTicketInfo) {
	v.TicketNo = ""
	v.GmtCreate = ""
	v.GmtModify = ""
	v.OriginTicketNo = ""
	v.ChangeFlightNo = ""
	v.ChangeCabin = ""
	v.ChangeCabinLevel = ""
	v.DepTime = ""
	v.ArrTime = ""
	v.ChangeReason = ""
	v.ChangeOrderId = 0
	v.ChangeType = 0
	v.ChangeFee = 0
	v.UpgradeFee = 0
	poolFlightChangeTicketInfo.Put(v)
}
