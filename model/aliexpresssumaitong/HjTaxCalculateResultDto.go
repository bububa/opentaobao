package aliexpresssumaitong

import (
	"sync"
)

// HjTaxCalculateResultDto 结构体
type HjTaxCalculateResultDto struct {
	// 交易行计税结果
	Lines []Lines `json:"lines,omitempty" xml:"lines>lines,omitempty"`
	// 币种转欧元汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 交易币种
	Currency *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
}

var poolHjTaxCalculateResultDto = sync.Pool{
	New: func() any {
		return new(HjTaxCalculateResultDto)
	},
}

// GetHjTaxCalculateResultDto() 从对象池中获取HjTaxCalculateResultDto
func GetHjTaxCalculateResultDto() *HjTaxCalculateResultDto {
	return poolHjTaxCalculateResultDto.Get().(*HjTaxCalculateResultDto)
}

// ReleaseHjTaxCalculateResultDto 释放HjTaxCalculateResultDto
func ReleaseHjTaxCalculateResultDto(v *HjTaxCalculateResultDto) {
	v.Lines = v.Lines[:0]
	v.ExchangeRate = ""
	v.Currency = nil
	poolHjTaxCalculateResultDto.Put(v)
}
