package btrip

// BtripHotelCancelPolicyDto 结构体
type BtripHotelCancelPolicyDto struct {
	// 详细的取消规则
	BtripHotelCancelPolicyInfoDTOList []BtripHotelCancelPolicyInfoDto `json:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" xml:"btrip_hotel_cancel_policy_info_d_t_o_list>btrip_hotel_cancel_policy_info_dto,omitempty"`
	// 取消类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}
