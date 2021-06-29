package maitix

// SeatQueryDTO 
type SeatQueryDTO struct {
    // 单票座位信息
    OrdinarySeatDTOS   []OrdinarySeatDTO `json:"ordinary_seat_d_t_o_s,omitempty" xml:"ordinary_seat_d_t_o_s>ordinary_seat_dto,omitempty"`
    // 座位状态，2有效
    SeatStatus   int64 `json:"seat_status,omitempty" xml:"seat_status,omitempty"`
    // 套票座位信息
    CombineSeatDTOS   []CombineSeatDTO `json:"combine_seat_d_t_o_s,omitempty" xml:"combine_seat_d_t_o_s>combine_seat_dto,omitempty"`
}
