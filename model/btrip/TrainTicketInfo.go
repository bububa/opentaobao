package btrip

import (
	"sync"
)

// TrainTicketInfo 结构体
type TrainTicketInfo struct {
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 车厢号
	CoachNo string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 座位类型
	SeatTypeName string `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 车次类型
	TrainTypeName string `json:"train_type_name,omitempty" xml:"train_type_name,omitempty"`
	// 出发时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 座位号
	SeatNo string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// 进站时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 出站时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 外部订单状态
	OutTicketStatus string `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	// 预订服务费
	ServiceFee float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 票价
	TicketPrice float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 支付方式 1个人支付/2企业支付/3混付
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 票状态
	TicketStatus int64 `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// 第几程，0：第一程，1：第二程
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 是否改签
	Changed bool `json:"changed,omitempty" xml:"changed,omitempty"`
}

var poolTrainTicketInfo = sync.Pool{
	New: func() any {
		return new(TrainTicketInfo)
	},
}

// GetTrainTicketInfo() 从对象池中获取TrainTicketInfo
func GetTrainTicketInfo() *TrainTicketInfo {
	return poolTrainTicketInfo.Get().(*TrainTicketInfo)
}

// ReleaseTrainTicketInfo 释放TrainTicketInfo
func ReleaseTrainTicketInfo(v *TrainTicketInfo) {
	v.TicketNo = ""
	v.CoachNo = ""
	v.GmtModify = ""
	v.SeatTypeName = ""
	v.GmtCreate = ""
	v.UserId = ""
	v.TrainTypeName = ""
	v.StartTime = ""
	v.EndTime = ""
	v.SeatNo = ""
	v.CheckInTime = ""
	v.CheckOutTime = ""
	v.OutTicketStatus = ""
	v.ServiceFee = 0
	v.TicketPrice = 0
	v.PayType = 0
	v.TicketStatus = 0
	v.SegmentIndex = 0
	v.Changed = false
	poolTrainTicketInfo.Put(v)
}
