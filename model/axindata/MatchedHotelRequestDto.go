package axindata

// MatchedHotelRequestDto 结构体
type MatchedHotelRequestDto struct {
	// 要匹配的酒店参数列表
	HotelMatchDTOList []HotelMatchDto `json:"hotel_match_d_t_o_list,omitempty" xml:"hotel_match_d_t_o_list>hotel_match_dto,omitempty"`
	// 分销商id
	DistributorTid int64 `json:"distributor_tid,omitempty" xml:"distributor_tid,omitempty"`
}
