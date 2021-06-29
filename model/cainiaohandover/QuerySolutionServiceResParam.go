package cainiaohandover

// QuerySolutionServiceResParam 
type QuerySolutionServiceResParam struct {
    // 解决方案code
    SolutionCode   string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
    // 服务参数
    ServiceParam   *ServiceParam `json:"service_param,omitempty" xml:"service_param,omitempty"`
}
