package flight

// CabinInfoDto 结构体
type CabinInfoDto struct {
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 乘机人类型，1成人2儿童3婴儿
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
