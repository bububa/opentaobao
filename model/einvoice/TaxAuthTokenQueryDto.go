package einvoice

import (
	"sync"
)

// TaxAuthTokenQueryDto 结构体
type TaxAuthTokenQueryDto struct {
	// 税控产品码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 服务商平台授权商户的税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
}

var poolTaxAuthTokenQueryDto = sync.Pool{
	New: func() any {
		return new(TaxAuthTokenQueryDto)
	},
}

// GetTaxAuthTokenQueryDto() 从对象池中获取TaxAuthTokenQueryDto
func GetTaxAuthTokenQueryDto() *TaxAuthTokenQueryDto {
	return poolTaxAuthTokenQueryDto.Get().(*TaxAuthTokenQueryDto)
}

// ReleaseTaxAuthTokenQueryDto 释放TaxAuthTokenQueryDto
func ReleaseTaxAuthTokenQueryDto(v *TaxAuthTokenQueryDto) {
	v.ProductCode = ""
	v.PayeeRegisterNo = ""
	poolTaxAuthTokenQueryDto.Put(v)
}
