package btrip

// HotelInfoRq 结构体
type HotelInfoRq struct {
	// 渠道子ID
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 可选项，自定义渠道名称（用于创建渠道）
	IsvName string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
