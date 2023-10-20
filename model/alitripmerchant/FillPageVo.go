package alitripmerchant

import (
	"sync"
)

// FillPageVo 结构体
type FillPageVo struct {
	// 订房必读
	ReservationNotice string `json:"reservation_notice,omitempty" xml:"reservation_notice,omitempty"`
	// 几点前离店
	CheckOutHour string `json:"check_out_hour,omitempty" xml:"check_out_hour,omitempty"`
	// 几点后入住
	CheckInHour string `json:"check_in_hour,omitempty" xml:"check_in_hour,omitempty"`
	// 星期几离店
	CheckOutWeek string `json:"check_out_week,omitempty" xml:"check_out_week,omitempty"`
	// 星期几入住
	CheckInWeek string `json:"check_in_week,omitempty" xml:"check_in_week,omitempty"`
	// 离店日期
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住日期
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 房间图片
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
	// 商品名称
	RpTitle string `json:"rp_title,omitempty" xml:"rp_title,omitempty"`
	// 房间窗型
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 房间床型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 房间大小
	RoomArea string `json:"room_area,omitempty" xml:"room_area,omitempty"`
	// 房间名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 酒店背景图片
	HotelUrl string `json:"hotel_url,omitempty" xml:"hotel_url,omitempty"`
	// 通用取消政策描述
	UniversalCancelPolicyDesc string `json:"universal_cancel_policy_desc,omitempty" xml:"universal_cancel_policy_desc,omitempty"`
	// 通用取消政策名称
	UniversalCancelPolicyName string `json:"universal_cancel_policy_name,omitempty" xml:"universal_cancel_policy_name,omitempty"`
	// 入住几晚
	TotalDay int64 `json:"total_day,omitempty" xml:"total_day,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
}

var poolFillPageVo = sync.Pool{
	New: func() any {
		return new(FillPageVo)
	},
}

// GetFillPageVo() 从对象池中获取FillPageVo
func GetFillPageVo() *FillPageVo {
	return poolFillPageVo.Get().(*FillPageVo)
}

// ReleaseFillPageVo 释放FillPageVo
func ReleaseFillPageVo(v *FillPageVo) {
	v.ReservationNotice = ""
	v.CheckOutHour = ""
	v.CheckInHour = ""
	v.CheckOutWeek = ""
	v.CheckInWeek = ""
	v.CheckOutDate = ""
	v.CheckInDate = ""
	v.FileUrl = ""
	v.RpTitle = ""
	v.WindowType = ""
	v.BedType = ""
	v.RoomArea = ""
	v.RoomName = ""
	v.Address = ""
	v.HotelName = ""
	v.HotelUrl = ""
	v.UniversalCancelPolicyDesc = ""
	v.UniversalCancelPolicyName = ""
	v.TotalDay = 0
	v.MaxOccupancy = 0
	poolFillPageVo.Put(v)
}
