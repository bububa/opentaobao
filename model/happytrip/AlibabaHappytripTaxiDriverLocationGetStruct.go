package happytrip

// AlibabaHappytripTaxiDriverLocationGetStruct 结构体
type AlibabaHappytripTaxiDriverLocationGetStruct struct {
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 司机方向，正北为0度，顺时针方向
	Direction string `json:"direction,omitempty" xml:"direction,omitempty"`
	// 车速，正数表示向前，负数表示向后，单位是 m/s
	Speed string `json:"speed,omitempty" xml:"speed,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
}
