package cainiaohandover

// OpenSolutionInquiryResponse 结构体
type OpenSolutionInquiryResponse struct {
	// 可用的解决方案列表
	UsableSolutionList []OpenSolutionDto `json:"usable_solution_list,omitempty" xml:"usable_solution_list>open_solution_dto,omitempty"`
}
