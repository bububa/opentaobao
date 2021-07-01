package cainiaohandover

// OpenSolutionDto 结构体
type OpenSolutionDto struct {
	// 时效信息
	TimingList []OpenTimingDto `json:"timing_list,omitempty" xml:"timing_list>open_timing_dto,omitempty"`
	// 费用列表
	FeeList []OpenFeeDto `json:"fee_list,omitempty" xml:"fee_list>open_fee_dto,omitempty"`
	// 推荐指数
	RecommendIndex int64 `json:"recommend_index,omitempty" xml:"recommend_index,omitempty"`
	// 解决方案code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 解决方案名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
