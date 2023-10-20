package cainiaohandover

import (
	"sync"
)

// QuerySolutionServiceResParam 结构体
type QuerySolutionServiceResParam struct {
	// 解决方案code
	SolutionCode string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
	// 服务参数
	ServiceParam *ServiceParam `json:"service_param,omitempty" xml:"service_param,omitempty"`
}

var poolQuerySolutionServiceResParam = sync.Pool{
	New: func() any {
		return new(QuerySolutionServiceResParam)
	},
}

// GetQuerySolutionServiceResParam() 从对象池中获取QuerySolutionServiceResParam
func GetQuerySolutionServiceResParam() *QuerySolutionServiceResParam {
	return poolQuerySolutionServiceResParam.Get().(*QuerySolutionServiceResParam)
}

// ReleaseQuerySolutionServiceResParam 释放QuerySolutionServiceResParam
func ReleaseQuerySolutionServiceResParam(v *QuerySolutionServiceResParam) {
	v.SolutionCode = ""
	v.ServiceParam = nil
	poolQuerySolutionServiceResParam.Put(v)
}
