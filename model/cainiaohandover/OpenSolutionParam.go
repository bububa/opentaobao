package cainiaohandover

import (
	"sync"
)

// OpenSolutionParam 结构体
type OpenSolutionParam struct {
	// 物流服务列表
	ServiceParams []OpenServiceParam `json:"service_params,omitempty" xml:"service_params>open_service_param,omitempty"`
	// 解决方案code
	SolutionCode string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
}

var poolOpenSolutionParam = sync.Pool{
	New: func() any {
		return new(OpenSolutionParam)
	},
}

// GetOpenSolutionParam() 从对象池中获取OpenSolutionParam
func GetOpenSolutionParam() *OpenSolutionParam {
	return poolOpenSolutionParam.Get().(*OpenSolutionParam)
}

// ReleaseOpenSolutionParam 释放OpenSolutionParam
func ReleaseOpenSolutionParam(v *OpenSolutionParam) {
	v.ServiceParams = v.ServiceParams[:0]
	v.SolutionCode = ""
	poolOpenSolutionParam.Put(v)
}
