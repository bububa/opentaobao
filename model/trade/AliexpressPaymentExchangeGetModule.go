package trade

import (
	"sync"
)

// AliexpressPaymentExchangeGetModule 结构体
type AliexpressPaymentExchangeGetModule struct {
	// 报价币种
	QuoteCurrency string `json:"quote_currency,omitempty" xml:"quote_currency,omitempty"`
	// 外部机构汇率号
	InstExchangeRateNo string `json:"inst_exchange_rate_no,omitempty" xml:"inst_exchange_rate_no,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 缓冲时间
	ThresholdTime string `json:"threshold_time,omitempty" xml:"threshold_time,omitempty"`
	// 汇率
	Rate string `json:"rate,omitempty" xml:"rate,omitempty"`
	// 有效时间
	ValidTime string `json:"valid_time,omitempty" xml:"valid_time,omitempty"`
	// 基准币种
	BaseCurrency string `json:"base_currency,omitempty" xml:"base_currency,omitempty"`
	// 汇率号
	ExchangeRateNo string `json:"exchange_rate_no,omitempty" xml:"exchange_rate_no,omitempty"`
	// 是否可交易
	Tradable bool `json:"tradable,omitempty" xml:"tradable,omitempty"`
}

var poolAliexpressPaymentExchangeGetModule = sync.Pool{
	New: func() any {
		return new(AliexpressPaymentExchangeGetModule)
	},
}

// GetAliexpressPaymentExchangeGetModule() 从对象池中获取AliexpressPaymentExchangeGetModule
func GetAliexpressPaymentExchangeGetModule() *AliexpressPaymentExchangeGetModule {
	return poolAliexpressPaymentExchangeGetModule.Get().(*AliexpressPaymentExchangeGetModule)
}

// ReleaseAliexpressPaymentExchangeGetModule 释放AliexpressPaymentExchangeGetModule
func ReleaseAliexpressPaymentExchangeGetModule(v *AliexpressPaymentExchangeGetModule) {
	v.QuoteCurrency = ""
	v.InstExchangeRateNo = ""
	v.ExpireTime = ""
	v.ThresholdTime = ""
	v.Rate = ""
	v.ValidTime = ""
	v.BaseCurrency = ""
	v.ExchangeRateNo = ""
	v.Tradable = false
	poolAliexpressPaymentExchangeGetModule.Put(v)
}
