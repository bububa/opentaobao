package cainiaohandover

import (
	"sync"
)

// SolutionServiceResQueryResponse 结构体
type SolutionServiceResQueryResponse struct {
	// 物流服务资源列表
	SolutionServiceResList []SolutionServiceResDto `json:"solution_service_res_list,omitempty" xml:"solution_service_res_list>solution_service_res_dto,omitempty"`
}

var poolSolutionServiceResQueryResponse = sync.Pool{
	New: func() any {
		return new(SolutionServiceResQueryResponse)
	},
}

// GetSolutionServiceResQueryResponse() 从对象池中获取SolutionServiceResQueryResponse
func GetSolutionServiceResQueryResponse() *SolutionServiceResQueryResponse {
	return poolSolutionServiceResQueryResponse.Get().(*SolutionServiceResQueryResponse)
}

// ReleaseSolutionServiceResQueryResponse 释放SolutionServiceResQueryResponse
func ReleaseSolutionServiceResQueryResponse(v *SolutionServiceResQueryResponse) {
	v.SolutionServiceResList = v.SolutionServiceResList[:0]
	poolSolutionServiceResQueryResponse.Put(v)
}
