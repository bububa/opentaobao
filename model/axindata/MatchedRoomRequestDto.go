package axindata

// MatchedRoomRequestDto 结构体
type MatchedRoomRequestDto struct {
	// 要匹配的房型参数列表
	HotelRoomMatchDTOList []HotelRoomMatchDto `json:"hotel_room_match_d_t_o_list,omitempty" xml:"hotel_room_match_d_t_o_list>hotel_room_match_dto,omitempty"`
	// 分销商淘宝账号id
	DistributorTid int64 `json:"distributor_tid,omitempty" xml:"distributor_tid,omitempty"`
}
