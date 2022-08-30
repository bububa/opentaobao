package btrip

// HotelShareInfo 结构体
type HotelShareInfo struct {
	// 合住公式。1-"(A+B)*param%",2-"A*param%",3-"A+B*param%",4-"A+param元",5-"(A+B)/2+param元"
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 请传入整数即可，当合住方式为1/2/3时接口会处理成x%
	Param string `json:"param,omitempty" xml:"param,omitempty"`
}
