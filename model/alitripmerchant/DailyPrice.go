package alitripmerchant

// DailyPrice 
type DailyPrice struct {
    // 价格
    ReallyPrice   string `json:"really_price,omitempty" xml:"really_price,omitempty"`
    // 日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
}
