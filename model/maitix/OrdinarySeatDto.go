package maitix

// OrdinarySeatDTO 
type OrdinarySeatDTO struct {
    // 座位号
    SeatNo   string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
    // 排号
    RowNo   string `json:"row_no,omitempty" xml:"row_no,omitempty"`
    // 楼层名称
    FloorName   string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
    // 看台名称
    StandName   string `json:"stand_name,omitempty" xml:"stand_name,omitempty"`
    // 看台ID
    StandId   int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
    // 座位ID
    SeatId   int64 `json:"seat_id,omitempty" xml:"seat_id,omitempty"`
    // 价格，单位分
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 票品ID,价格id,和ticket_item_id等价
    PriceId   int64 `json:"price_id,omitempty" xml:"price_id,omitempty"`
}
