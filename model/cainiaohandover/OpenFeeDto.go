package cainiaohandover

// OpenFeeDto 结构体
type OpenFeeDto struct {
	// 总费用
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 费用类型，POST_ESTIMATED_COST：预估费用
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 费用详细列表
	FeeDetailList []OpenFeeDetailDto `json:"fee_detail_list,omitempty" xml:"fee_detail_list>open_fee_detail_dto,omitempty"`
}
