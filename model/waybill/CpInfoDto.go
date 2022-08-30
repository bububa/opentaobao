package waybill

// CpInfoDto 结构体
type CpInfoDto struct {
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// cp名称
	CpName string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
}
