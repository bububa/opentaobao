package simba

// AdzoneConfigVo 结构体
type AdzoneConfigVo struct {
	// 描述
	AimDesc string `json:"aim_desc,omitempty" xml:"aim_desc,omitempty"`
	// 资源包id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 是否支持溢价
	Discount bool `json:"discount,omitempty" xml:"discount,omitempty"`
}
