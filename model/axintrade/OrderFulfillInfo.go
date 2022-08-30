package axintrade

// OrderFulfillInfo 结构体
type OrderFulfillInfo struct {
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 最晚到店时间
	LateArriveTime string `json:"late_arrive_time,omitempty" xml:"late_arrive_time,omitempty"`
	// 酒店确认号
	ConfirmCode string `json:"confirm_code,omitempty" xml:"confirm_code,omitempty"`
	// 房间数量
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
}
