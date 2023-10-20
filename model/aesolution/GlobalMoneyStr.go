package aesolution

import (
	"sync"
)

// GlobalMoneyStr 结构体
type GlobalMoneyStr struct {
	// Currency code
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Amount
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// Amount
	AmountStr string `json:"amount_str,omitempty" xml:"amount_str,omitempty"`
	// Cent
	Cent int64 `json:"cent,omitempty" xml:"cent,omitempty"`
	// The factor to the smallest currency unit
	CentFactor int64 `json:"cent_factor,omitempty" xml:"cent_factor,omitempty"`
	// Currency description
	Currency *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
}

var poolGlobalMoneyStr = sync.Pool{
	New: func() any {
		return new(GlobalMoneyStr)
	},
}

// GetGlobalMoneyStr() 从对象池中获取GlobalMoneyStr
func GetGlobalMoneyStr() *GlobalMoneyStr {
	return poolGlobalMoneyStr.Get().(*GlobalMoneyStr)
}

// ReleaseGlobalMoneyStr 释放GlobalMoneyStr
func ReleaseGlobalMoneyStr(v *GlobalMoneyStr) {
	v.CurrencyCode = ""
	v.Amount = ""
	v.AmountStr = ""
	v.Cent = 0
	v.CentFactor = 0
	v.Currency = nil
	poolGlobalMoneyStr.Put(v)
}
