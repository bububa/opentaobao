package hotel

// SRoomTypeDailyPrice 
type SRoomTypeDailyPrice struct {

    // 离店日期
    
    End   string `json:"end,omitempty" xml:"end,omitempty"`
    

    // 最低价
    
    LowPrice   int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
    

    // 当前标准房型下所有库价集合
    
    RoomTypeDailyPriceList   []RoomTypeDailyPrice `json:"room_type_daily_price_list,omitempty" xml:"room_type_daily_price_list,omitempty"`
    

    // 标准酒店id
    
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    

    // 标准房型id
    
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    

    // 入住时间
    
    Start   string `json:"start,omitempty" xml:"start,omitempty"`
    

    // 床型字符串
    
    BedTypeString   string `json:"bed_type_string,omitempty" xml:"bed_type_string,omitempty"`
    

    // 标准房型名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 窗型枚举，0-无窗,1-有窗,2-部分有窗,3-暗窗,4-部分暗窗
    
    WindowType   string `json:"window_type,omitempty" xml:"window_type,omitempty"`
    

}
