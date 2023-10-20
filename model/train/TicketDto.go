package train

import (
	"sync"
)

// TicketDto 结构体
type TicketDto struct {
	// 12306订单号
	SequenceNo string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 座席号
	SeatNo string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// 车厢号
	CoachNo string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 票类型（1 成人票 2 儿童 3 学生 4 残、军）
	TicketTypeCode string `json:"ticket_type_code,omitempty" xml:"ticket_type_code,omitempty"`
	// 真实座席编码
	RealSeatTypeCode string `json:"real_seat_type_code,omitempty" xml:"real_seat_type_code,omitempty"`
	// 出发站名称
	FromStationName string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// 中转站编码
	InterChangeStationTelecode string `json:"inter_change_station_telecode,omitempty" xml:"inter_change_station_telecode,omitempty"`
	// 中转站名称
	InterChangeStationName string `json:"inter_change_station_name,omitempty" xml:"inter_change_station_name,omitempty"`
	// 车次编码
	TrainCode string `json:"train_code,omitempty" xml:"train_code,omitempty"`
	// 坐席编码
	SeatTypeCode string `json:"seat_type_code,omitempty" xml:"seat_type_code,omitempty"`
	// 座席名称
	SeatTypeName string `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	// 到达站名称
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// 出发站编码
	FromStationTelecode string `json:"from_station_telecode,omitempty" xml:"from_station_telecode,omitempty"`
	// 车次日期
	TrainDate string `json:"train_date,omitempty" xml:"train_date,omitempty"`
	// 出发时间
	FromTime string `json:"from_time,omitempty" xml:"from_time,omitempty"`
	// 到达站编码
	ToStationTelecode string `json:"to_station_telecode,omitempty" xml:"to_station_telecode,omitempty"`
	// 到达时间
	ToTime string `json:"to_time,omitempty" xml:"to_time,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 子单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// ttp子单号
	TtpSubOrderId int64 `json:"ttp_sub_order_id,omitempty" xml:"ttp_sub_order_id,omitempty"`
	// 真实票价 单位 分
	RealTicketPrice int64 `json:"real_ticket_price,omitempty" xml:"real_ticket_price,omitempty"`
	// 定制票类型(0:不指定坐席 1:下铺 2:下铺or中铺 3:过道 4:靠窗)
	VipCustomType int64 `json:"vip_custom_type,omitempty" xml:"vip_custom_type,omitempty"`
	// 程ID
	SegmentId int64 `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// 程 序号
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

var poolTicketDto = sync.Pool{
	New: func() any {
		return new(TicketDto)
	},
}

// GetTicketDto() 从对象池中获取TicketDto
func GetTicketDto() *TicketDto {
	return poolTicketDto.Get().(*TicketDto)
}

// ReleaseTicketDto 释放TicketDto
func ReleaseTicketDto(v *TicketDto) {
	v.SequenceNo = ""
	v.SeatNo = ""
	v.CoachNo = ""
	v.BatchNo = ""
	v.TicketTypeCode = ""
	v.RealSeatTypeCode = ""
	v.FromStationName = ""
	v.InterChangeStationTelecode = ""
	v.InterChangeStationName = ""
	v.TrainCode = ""
	v.SeatTypeCode = ""
	v.SeatTypeName = ""
	v.ToStationName = ""
	v.FromStationTelecode = ""
	v.TrainDate = ""
	v.FromTime = ""
	v.ToStationTelecode = ""
	v.ToTime = ""
	v.TtpOrderId = 0
	v.SubOrderId = 0
	v.TtpSubOrderId = 0
	v.RealTicketPrice = 0
	v.VipCustomType = 0
	v.SegmentId = 0
	v.SegmentIndex = 0
	poolTicketDto.Put(v)
}
