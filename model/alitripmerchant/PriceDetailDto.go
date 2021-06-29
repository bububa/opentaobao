package alitripmerchant

// PriceDetailDTO 
type PriceDetailDTO struct {
    // 总价格
    TotalPrice   string `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 总税费
    TotalTax   string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
    // 总房间价格(不含税)
    TotalRoomPrice   string `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
    // 每日价格
    DailyPrices   []DailyPrice `json:"daily_prices,omitempty" xml:"daily_prices>daily_price,omitempty"`
    // 货币
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 入住时间
    Days   int64 `json:"days,omitempty" xml:"days,omitempty"`
    // 房间数量
    RoomNum   int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
}
