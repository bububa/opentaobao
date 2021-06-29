package xhotelonlineorder

// DailyRoomFee 
type DailyRoomFee struct {
    // 无
    DailyPrices   []DailyPrice `json:"daily_prices,omitempty" xml:"daily_prices>daily_price,omitempty"`
}
