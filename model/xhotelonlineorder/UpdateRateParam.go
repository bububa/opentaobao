package xhotelonlineorder

// UpdateRateParam 结构体
type UpdateRateParam struct {
	// rate更新列表
	UpdateRateDOList []UpdateRateDo `json:"update_rate_d_o_list,omitempty" xml:"update_rate_d_o_list>update_rate_do,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 供应商
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
}
