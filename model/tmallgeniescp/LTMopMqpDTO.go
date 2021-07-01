package tmallgeniescp

// LTMopMqpDto 结构体
type LTMopMqpDto struct {
	// 物料编码
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 扩展字段
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// mpq
	Mpq string `json:"mpq,omitempty" xml:"mpq,omitempty"`
	// moq
	Moq string `json:"moq,omitempty" xml:"moq,omitempty"`
	// leadtime
	LeadTime string `json:"lead_time,omitempty" xml:"lead_time,omitempty"`
	// 供应商编码
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
}
