package axindata

// BedGroupDto 结构体
type BedGroupDto struct {
	// 床型详细信息
	BedInfoDTOList []BedInfoDto `json:"bed_info_d_t_o_list,omitempty" xml:"bed_info_d_t_o_list>bed_info_dto,omitempty"`
}
