package axindata

// HotelCityVo 结构体
type HotelCityVo struct {
	// 城市名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 城市编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
