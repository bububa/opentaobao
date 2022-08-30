package axintrade

// TicketInfoDto 结构体
type TicketInfoDto struct {
	// 景区名称
	ScenicName string `json:"scenic_name,omitempty" xml:"scenic_name,omitempty"`
	// 收费项目名称
	SpuName string `json:"spu_name,omitempty" xml:"spu_name,omitempty"`
	// 票种名称
	TicketKindName string `json:"ticket_kind_name,omitempty" xml:"ticket_kind_name,omitempty"`
	// 场次名称
	SessionName string `json:"session_name,omitempty" xml:"session_name,omitempty"`
	// 区域名称
	RegionName string `json:"region_name,omitempty" xml:"region_name,omitempty"`
}
