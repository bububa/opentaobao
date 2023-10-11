package usergrowth

// Temperature 结构体
type Temperature struct {
	// 当日最高温度
	Highest string `json:"highest,omitempty" xml:"highest,omitempty"`
	// 当日最低温度
	Lowest string `json:"lowest,omitempty" xml:"lowest,omitempty"`
}
