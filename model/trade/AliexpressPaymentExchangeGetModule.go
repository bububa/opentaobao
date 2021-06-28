package trade

// AliexpressPaymentExchangeGetModule 
/* model for simplify = false
type AliexpressPaymentExchangeGetModule struct {

    // 报价币种
    
    QuoteCurrency   string `json:"quote_currency,omitempty"`
    

    // 外部机构汇率号
    
    InstExchangeRateNo   string `json:"inst_exchange_rate_no,omitempty"`
    

    // 过期时间
    
    ExpireTime   string `json:"expire_time,omitempty"`
    

    // 缓冲时间
    
    ThresholdTime   string `json:"threshold_time,omitempty"`
    

    // 汇率
    
    Rate   string `json:"rate,omitempty"`
    

    // 是否可交易
    
    Tradable   bool `json:"tradable,omitempty"`
    

    // 有效时间
    
    ValidTime   string `json:"valid_time,omitempty"`
    

    // 基准币种
    
    BaseCurrency   string `json:"base_currency,omitempty"`
    

    // 汇率号
    
    ExchangeRateNo   string `json:"exchange_rate_no,omitempty"`
    

}
*/

// AliexpressPaymentExchangeGetModule 
type AliexpressPaymentExchangeGetModule struct {

    // 报价币种
    QuoteCurrency   string `json:"quote_currency,omitempty"`

    // 外部机构汇率号
    InstExchangeRateNo   string `json:"inst_exchange_rate_no,omitempty"`

    // 过期时间
    ExpireTime   string `json:"expire_time,omitempty"`

    // 缓冲时间
    ThresholdTime   string `json:"threshold_time,omitempty"`

    // 汇率
    Rate   string `json:"rate,omitempty"`

    // 是否可交易
    Tradable   bool `json:"tradable,omitempty"`

    // 有效时间
    ValidTime   string `json:"valid_time,omitempty"`

    // 基准币种
    BaseCurrency   string `json:"base_currency,omitempty"`

    // 汇率号
    ExchangeRateNo   string `json:"exchange_rate_no,omitempty"`

}
