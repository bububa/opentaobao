package axintrade

// PackageHotelPolicyDto 结构体
type PackageHotelPolicyDto struct {
	// 入住方式
	CheckInTypes []string `json:"check_in_types,omitempty" xml:"check_in_types>string,omitempty"`
	// 入住时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 离店时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 入住方式其他说明
	CheckInTypeRemark string `json:"check_in_type_remark,omitempty" xml:"check_in_type_remark,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}
