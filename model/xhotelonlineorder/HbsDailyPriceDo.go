package xhotelonlineorder

// HbsDailyPriceDo 
type HbsDailyPriceDo struct {

    // 币种
    
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    

    // 汇率
    
    ExchangeRate   string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
    

    // 每日价格
    
    DailyPrices   []DailyPriceTo `json:"daily_prices,omitempty" xml:"daily_prices,omitempty"`
    

}
