package maitix

// StandQueryParam 
type StandQueryParam struct {
    // 座位ID
    SeatIds   []int64 `json:"seat_ids,omitempty" xml:"seat_ids>int64,omitempty"`
    // 看台ID
    StandId   int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
}
