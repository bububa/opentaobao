package hotel

// HotelTopStaticDetailsParam 结构体
type HotelTopStaticDetailsParam struct {
	// 城市code
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 渠道跟踪号
	Ttid string `json:"ttid,omitempty" xml:"ttid,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
