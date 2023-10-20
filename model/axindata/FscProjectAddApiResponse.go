package axindata

import (
	"sync"
)

// FscProjectAddApiResponse 结构体
type FscProjectAddApiResponse struct {
	// 团期计划ID映射
	ProjectMapDTOS []FscProjectMapDto `json:"project_map_d_t_o_s,omitempty" xml:"project_map_d_t_o_s>fsc_project_map_dto,omitempty"`
}

var poolFscProjectAddApiResponse = sync.Pool{
	New: func() any {
		return new(FscProjectAddApiResponse)
	},
}

// GetFscProjectAddApiResponse() 从对象池中获取FscProjectAddApiResponse
func GetFscProjectAddApiResponse() *FscProjectAddApiResponse {
	return poolFscProjectAddApiResponse.Get().(*FscProjectAddApiResponse)
}

// ReleaseFscProjectAddApiResponse 释放FscProjectAddApiResponse
func ReleaseFscProjectAddApiResponse(v *FscProjectAddApiResponse) {
	v.ProjectMapDTOS = v.ProjectMapDTOS[:0]
	poolFscProjectAddApiResponse.Put(v)
}
