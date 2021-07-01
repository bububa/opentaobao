package ieagency

// IeBaggageVo 结构体
type IeBaggageVo struct {
	// 行李总件数
	Pc int64 `json:"pc,omitempty" xml:"pc,omitempty"`
	// 行李重量，单位KG，当allWeight为是的时候行李重量描述所有件数的总重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 是否所有行李重量
	IsAllWeight bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
}
