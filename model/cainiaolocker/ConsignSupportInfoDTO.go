package cainiaolocker

// ConsignSupportInfoDto 结构体
type ConsignSupportInfoDto struct {
	// 支持CP的code列表
	SupportCpList []string `json:"support_cp_list,omitempty" xml:"support_cp_list>string,omitempty"`
}
