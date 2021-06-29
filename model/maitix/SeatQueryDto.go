package maitix

// SeatQueryDto 
type SeatQueryDto struct {

    // 单票座位信息
    
    OrdinarySeatDTOS   []OrdinarySeatDto `json:"ordinary_seat_d_t_o_s,omitempty" xml:"ordinary_seat_d_t_o_s,omitempty"`
    

    // 座位状态，2有效
    
    SeatStatus   int64 `json:"seat_status,omitempty" xml:"seat_status,omitempty"`
    

    // 套票座位信息
    
    CombineSeatDTOS   []CombineSeatDto `json:"combine_seat_d_t_o_s,omitempty" xml:"combine_seat_d_t_o_s,omitempty"`
    

}
