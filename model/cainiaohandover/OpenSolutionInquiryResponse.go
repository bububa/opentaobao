package cainiaohandover

import (
	"sync"
)

// OpenSolutionInquiryResponse 结构体
type OpenSolutionInquiryResponse struct {
	// 可用的解决方案列表
	UsableSolutionList []OpenSolutionDto `json:"usable_solution_list,omitempty" xml:"usable_solution_list>open_solution_dto,omitempty"`
}

var poolOpenSolutionInquiryResponse = sync.Pool{
	New: func() any {
		return new(OpenSolutionInquiryResponse)
	},
}

// GetOpenSolutionInquiryResponse() 从对象池中获取OpenSolutionInquiryResponse
func GetOpenSolutionInquiryResponse() *OpenSolutionInquiryResponse {
	return poolOpenSolutionInquiryResponse.Get().(*OpenSolutionInquiryResponse)
}

// ReleaseOpenSolutionInquiryResponse 释放OpenSolutionInquiryResponse
func ReleaseOpenSolutionInquiryResponse(v *OpenSolutionInquiryResponse) {
	v.UsableSolutionList = v.UsableSolutionList[:0]
	poolOpenSolutionInquiryResponse.Put(v)
}
