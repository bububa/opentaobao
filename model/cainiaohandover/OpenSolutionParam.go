package cainiaohandover

// OpenSolutionParam 
type OpenSolutionParam struct {

    // 解决方案code
    
    SolutionCode   string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
    

    // 物流服务列表
    
    ServiceParams   []OpenServiceParam `json:"service_params,omitempty" xml:"service_params,omitempty"`
    

}
