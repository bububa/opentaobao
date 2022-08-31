package btrip

// HotelShareInfo 结构体
type HotelShareInfo struct {
	// 合住公式。1-&#34;(A+B)*param%&#34;,2-&#34;A*param%&#34;,3-&#34;A+B*param%&#34;,4-&#34;A+param元&#34;,5-&#34;(A+B)/2+param元&#34;
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 请传入整数即可，当合住方式为1/2/3时接口会处理成x%
	Param string `json:"param,omitempty" xml:"param,omitempty"`
}
