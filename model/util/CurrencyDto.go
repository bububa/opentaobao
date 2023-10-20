package util

import (
	"sync"
)

// CurrencyDto 结构体
type CurrencyDto struct {
	// 货币编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 货币显示标示
	Symbol string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// 货币名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolCurrencyDto = sync.Pool{
	New: func() any {
		return new(CurrencyDto)
	},
}

// GetCurrencyDto() 从对象池中获取CurrencyDto
func GetCurrencyDto() *CurrencyDto {
	return poolCurrencyDto.Get().(*CurrencyDto)
}

// ReleaseCurrencyDto 释放CurrencyDto
func ReleaseCurrencyDto(v *CurrencyDto) {
	v.Code = ""
	v.Symbol = ""
	v.Name = ""
	poolCurrencyDto.Put(v)
}
