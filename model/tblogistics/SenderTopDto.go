package tblogistics

// SenderTopDto 结构体
type SenderTopDto struct {
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人电话，支持手机、座机
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 地址门牌号
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度（高德）
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度（高德）
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
}
