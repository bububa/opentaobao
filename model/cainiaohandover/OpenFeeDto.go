package cainiaohandover

import (
	"sync"
)

// OpenFeeDto 结构体
type OpenFeeDto struct {
	// 费用详细列表
	FeeDetailList []OpenFeeDetailDto `json:"fee_detail_list,omitempty" xml:"fee_detail_list>open_fee_detail_dto,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 费用类型，POST_ESTIMATED_COST：预估费用
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 总费用
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
}

var poolOpenFeeDto = sync.Pool{
	New: func() any {
		return new(OpenFeeDto)
	},
}

// GetOpenFeeDto() 从对象池中获取OpenFeeDto
func GetOpenFeeDto() *OpenFeeDto {
	return poolOpenFeeDto.Get().(*OpenFeeDto)
}

// ReleaseOpenFeeDto 释放OpenFeeDto
func ReleaseOpenFeeDto(v *OpenFeeDto) {
	v.FeeDetailList = v.FeeDetailList[:0]
	v.Currency = ""
	v.FeeType = ""
	v.TotalFee = 0
	poolOpenFeeDto.Put(v)
}
