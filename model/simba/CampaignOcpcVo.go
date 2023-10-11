package simba

// CampaignOcpcVo 结构体
type CampaignOcpcVo struct {
	// OCPC是否开启，0:否，1:是
	EnableOcpc string `json:"enable_ocpc,omitempty" xml:"enable_ocpc,omitempty"`
	// OCPC溢价比例
	OcpcRatio int64 `json:"ocpc_ratio,omitempty" xml:"ocpc_ratio,omitempty"`
}
