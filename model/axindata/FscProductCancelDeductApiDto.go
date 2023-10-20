package axindata

import (
	"sync"
)

// FscProductCancelDeductApiDto 结构体
type FscProductCancelDeductApiDto struct {
	// 责任方类型 DISTRIBUTOR：分销商 SUPPLIER：供应商
	ResponsibleType string `json:"responsible_type,omitempty" xml:"responsible_type,omitempty"`
	// 扣除金额类型 PERCENT：百分比 ABSOLUTE：绝对值
	DeductType string `json:"deduct_type,omitempty" xml:"deduct_type,omitempty"`
	// 扣除金额值 百分比例如：6.8代表扣除6.8% 绝对值例如：1.23代表扣除1.23元
	DeductValue string `json:"deduct_value,omitempty" xml:"deduct_value,omitempty"`
}

var poolFscProductCancelDeductApiDto = sync.Pool{
	New: func() any {
		return new(FscProductCancelDeductApiDto)
	},
}

// GetFscProductCancelDeductApiDto() 从对象池中获取FscProductCancelDeductApiDto
func GetFscProductCancelDeductApiDto() *FscProductCancelDeductApiDto {
	return poolFscProductCancelDeductApiDto.Get().(*FscProductCancelDeductApiDto)
}

// ReleaseFscProductCancelDeductApiDto 释放FscProductCancelDeductApiDto
func ReleaseFscProductCancelDeductApiDto(v *FscProductCancelDeductApiDto) {
	v.ResponsibleType = ""
	v.DeductType = ""
	v.DeductValue = ""
	poolFscProductCancelDeductApiDto.Put(v)
}
