package train

import (
	"sync"
)

// TapSubOrderVo 结构体
type TapSubOrderVo struct {
	// 出发站
	FromStationName string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// 出发站三字码
	FromStationTelecode string `json:"from_station_telecode,omitempty" xml:"from_station_telecode,omitempty"`
	// 到达站
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// 到达站三字码
	ToStationTelecode string `json:"to_station_telecode,omitempty" xml:"to_station_telecode,omitempty"`
	// 发车日期
	TrainDate string `json:"train_date,omitempty" xml:"train_date,omitempty"`
	// 车次编号
	TrainCode string `json:"train_code,omitempty" xml:"train_code,omitempty"`
	// 座席编号
	SeatTypeCode string `json:"seat_type_code,omitempty" xml:"seat_type_code,omitempty"`
	// 座席名称
	SeatTypeName string `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	// 在线选座参数(GDC支持在线选座，此坐席为优选坐席，出票员可以优先保证此字段解析的坐席)
	OnlineBookSeat string `json:"online_book_seat,omitempty" xml:"online_book_seat,omitempty"`
	// 中转站三字码
	InterChangeStationTelecode string `json:"inter_change_station_telecode,omitempty" xml:"inter_change_station_telecode,omitempty"`
	// 中转站名称
	InterChangeStationName string `json:"inter_change_station_name,omitempty" xml:"inter_change_station_name,omitempty"`
	// 出发时间
	FromTime string `json:"from_time,omitempty" xml:"from_time,omitempty"`
	// 到达时间
	ToTime string `json:"to_time,omitempty" xml:"to_time,omitempty"`
	// 子单状态
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 子单出票截止时间（yyyy-MM-dd HH:mm:ss）
	LastIssueTime string `json:"last_issue_time,omitempty" xml:"last_issue_time,omitempty"`
	// 主订单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 票价（单位 分）
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 定制票类型（0:不指定坐席 1:下铺 2:下铺or中铺 3:过道 4:靠窗）
	VipCustomType int64 `json:"vip_custom_type,omitempty" xml:"vip_custom_type,omitempty"`
	// 程id
	SegmentId int64 `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// 程 序号（区分中转一、二程）
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 是否为紧急单
	Emergency bool `json:"emergency,omitempty" xml:"emergency,omitempty"`
}

var poolTapSubOrderVo = sync.Pool{
	New: func() any {
		return new(TapSubOrderVo)
	},
}

// GetTapSubOrderVo() 从对象池中获取TapSubOrderVo
func GetTapSubOrderVo() *TapSubOrderVo {
	return poolTapSubOrderVo.Get().(*TapSubOrderVo)
}

// ReleaseTapSubOrderVo 释放TapSubOrderVo
func ReleaseTapSubOrderVo(v *TapSubOrderVo) {
	v.FromStationName = ""
	v.FromStationTelecode = ""
	v.ToStationName = ""
	v.ToStationTelecode = ""
	v.TrainDate = ""
	v.TrainCode = ""
	v.SeatTypeCode = ""
	v.SeatTypeName = ""
	v.OnlineBookSeat = ""
	v.InterChangeStationTelecode = ""
	v.InterChangeStationName = ""
	v.FromTime = ""
	v.ToTime = ""
	v.StatusName = ""
	v.LastIssueTime = ""
	v.TtpOrderId = 0
	v.SubOrderId = 0
	v.TicketPrice = 0
	v.VipCustomType = 0
	v.SegmentId = 0
	v.SegmentIndex = 0
	v.Emergency = false
	poolTapSubOrderVo.Put(v)
}
