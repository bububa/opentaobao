package cainiaohandover

// SolutionServiceResQueryResponse 结构体
type SolutionServiceResQueryResponse struct {
	// 物流服务资源列表
	SolutionServiceResList []SolutionServiceResDto `json:"solution_service_res_list,omitempty" xml:"solution_service_res_list>solution_service_res_dto,omitempty"`
}
