package hotel

// SHotelPrice 
/* model for simplify = false
type SHotelPrice struct {

    // 每个标准酒店某一天的所有库价集合
    
    DailyPriceList  struct {
        SHotelDailyPrice  []SHotelDailyPrice `json:"s_hotel_daily_price,omitempty"`
    } `json:"daily_price_list,omitempty"`
    

    // 标准房型shid
    
    Shid   int64 `json:"shid,omitempty"`
    

}
*/

// SHotelPrice 
type SHotelPrice struct {

    // 每个标准酒店某一天的所有库价集合
    DailyPriceList   []SHotelDailyPrice `json:"daily_price_list,omitempty"`

    // 标准房型shid
    Shid   int64 `json:"shid,omitempty"`

}
