package btrip

// HotelSearchListRq 结构体
type HotelSearchListRq struct {
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 城市code
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 排序方式，0-默认|1-销量|2-价格|3-距离|4-好评
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 排序方向，1-升序 or 0-降序
	Dir int64 `json:"dir,omitempty" xml:"dir,omitempty"`
	// 酒店id，多个用英文逗号分隔符
	Shids string `json:"shids,omitempty" xml:"shids,omitempty"`
}
