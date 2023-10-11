package tblogistics

// LocationTopDto 结构体
type LocationTopDto struct {
	// 省/直辖市
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 地级市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区/县
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 乡/镇/街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 经度，高德地图
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度，高德地图
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
}
