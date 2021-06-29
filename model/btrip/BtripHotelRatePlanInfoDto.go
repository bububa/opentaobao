package btrip

// BtripHotelRatePlanInfoDTO 
type BtripHotelRatePlanInfoDTO struct {
    // 房型描述
    BedDesc   string `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
    // 取消政策
    BtripHotelCancelPolicyDTO   *BtripHotelCancelPolicyDTO `json:"btrip_hotel_cancel_policy_d_t_o,omitempty" xml:"btrip_hotel_cancel_policy_d_t_o,omitempty"`
    // 最早入住时间
    EarliestCheckInTime   string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
    // 最晚离店时间
    LatestCheckOutTime   string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
    // 最大可预订房间数
    MaxBookingNum   int64 `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
    // 状态库存场景，不生效。值大于0场景使用
    MaxInventory   int64 `json:"max_inventory,omitempty" xml:"max_inventory,omitempty"`
    // 每间房最大可入住人数
    MaxOccupancyNum   int64 `json:"max_occupancy_num,omitempty" xml:"max_occupancy_num,omitempty"`
    // 是否需要填写电子邮箱
    NeedEmail   bool `json:"need_email,omitempty" xml:"need_email,omitempty"`
    // 每间房rate信息
    RateUnits   []BtripHotelRateUnitDTO `json:"rate_units,omitempty" xml:"rate_units>btrip_hotel_rate_unit_dto,omitempty"`
    // 总房价（会员价）
    TotalMemberRoomPrice   int64 `json:"total_member_room_price,omitempty" xml:"total_member_room_price,omitempty"`
    // 总房价
    TotalRoomPrice   int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
}
