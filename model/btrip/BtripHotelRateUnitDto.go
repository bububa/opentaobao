package btrip

// BtripHotelRateUnitDto 
type BtripHotelRateUnitDto struct {

    // 日历信息
    
    DailyPriceInfoList   []BtripHotelDailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list,omitempty"`
    

    // 最小售卖单元唯一key
    
    RateKey   string `json:"rate_key,omitempty" xml:"rate_key,omitempty"`
    

}
