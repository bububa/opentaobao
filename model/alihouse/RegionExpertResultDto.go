package alihouse

// RegionExpertResultDto 结构体
type RegionExpertResultDto struct {
	// 外部经纪人ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 异常信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
