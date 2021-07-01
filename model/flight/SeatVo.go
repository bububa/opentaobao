package flight

// SeatVo 结构体
type SeatVo struct {
	// 座位列 1:靠过道，2:靠窗，3:并排
	SeatArea int64 `json:"seat_area,omitempty" xml:"seat_area,omitempty"`
	// 排次 1:前排，2:中排，3:后排
	SeatRow int64 `json:"seat_row,omitempty" xml:"seat_row,omitempty"`
}
