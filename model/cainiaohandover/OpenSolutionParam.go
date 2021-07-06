package cainiaohandover

// OpenSolutionParam 结构体
type OpenSolutionParam struct {
	// 物流服务列表
	ServiceParams []OpenServiceParam `json:"service_params,omitempty" xml:"service_params>open_service_param,omitempty"`
	// 解决方案code
	SolutionCode string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
}
