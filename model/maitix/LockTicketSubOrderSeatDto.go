package maitix

// LockTicketSubOrderSeatDto 
type LockTicketSubOrderSeatDto struct {
    // 项目ID
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 项目名称
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    // 场次ID
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 场次名称
    PerformName   string `json:"perform_name,omitempty" xml:"perform_name,omitempty"`
    // 票品ID
    PriceId   int64 `json:"price_id,omitempty" xml:"price_id,omitempty"`
    // 票品名称
    PriceName   string `json:"price_name,omitempty" xml:"price_name,omitempty"`
    // 入场口
    Entry   string `json:"entry,omitempty" xml:"entry,omitempty"`
    // 看台ID
    StandId   int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
    // 看台名称
    StandName   string `json:"stand_name,omitempty" xml:"stand_name,omitempty"`
    // 楼层ID
    SeatFloorId   int64 `json:"seat_floor_id,omitempty" xml:"seat_floor_id,omitempty"`
    // 楼层名称
    SeatFloorName   string `json:"seat_floor_name,omitempty" xml:"seat_floor_name,omitempty"`
    // 区域ID
    SeatAreaId   int64 `json:"seat_area_id,omitempty" xml:"seat_area_id,omitempty"`
    // 区域名称
    SeatAreaName   string `json:"seat_area_name,omitempty" xml:"seat_area_name,omitempty"`
    // 座位分组，0:无座 1:有座
    SeatGroup   int64 `json:"seat_group,omitempty" xml:"seat_group,omitempty"`
    // 套票ID
    CombineId   string `json:"combine_id,omitempty" xml:"combine_id,omitempty"`
    // 座位ID
    SeatId   int64 `json:"seat_id,omitempty" xml:"seat_id,omitempty"`
    // 座位名称
    SeatName   string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
    // 座位排ID
    SeatRowId   int64 `json:"seat_row_id,omitempty" xml:"seat_row_id,omitempty"`
    // 座位排名称
    SeatRowName   string `json:"seat_row_name,omitempty" xml:"seat_row_name,omitempty"`
    // 座位类型
    SeatType   *model.File `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
}
