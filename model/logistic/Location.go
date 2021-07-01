package logistic

// Location 结构体
type Location struct {
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 邮政编码
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 详细地址，最大256个字节（128个中文）
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 所在城市（中文名称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 所在省份（中文名称）
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 国家名称
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 区/县（只适用于物流API）
	District string `json:"district,omitempty" xml:"district,omitempty"`
}
