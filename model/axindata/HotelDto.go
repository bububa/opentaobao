package axindata

// HotelDto 结构体
type HotelDto struct {
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
}
