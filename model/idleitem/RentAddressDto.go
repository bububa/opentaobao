package idleitem

// RentAddressDto 结构体
type RentAddressDto struct {
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 详细地址
	FullAddr string `json:"full_addr,omitempty" xml:"full_addr,omitempty"`
}
