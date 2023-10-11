package axindata

// FscSaleCommissionApiDto 结构体
type FscSaleCommissionApiDto struct {
	// 佣金类型
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 佣金数值
	CommissionNum string `json:"commission_num,omitempty" xml:"commission_num,omitempty"`
}
