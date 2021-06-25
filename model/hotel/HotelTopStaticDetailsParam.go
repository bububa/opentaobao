package hotel

// HotelTopStaticDetailsParam 
type HotelTopStaticDetailsParam struct {

    // 城市code
    CityCode   string `json:"city_code,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // 渠道跟踪号
    Ttid   string `json:"ttid,omitempty"`

}
