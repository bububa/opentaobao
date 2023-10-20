package tmallgeniescp

import (
	"sync"
)

// ConsensusDemandRequest 结构体
type ConsensusDemandRequest struct {
	// 入参对象列表
	ConsensusDemandParamDTOList []ConsensusDemandParamDto `json:"consensus_demand_param_d_t_o_list,omitempty" xml:"consensus_demand_param_d_t_o_list>consensus_demand_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolConsensusDemandRequest = sync.Pool{
	New: func() any {
		return new(ConsensusDemandRequest)
	},
}

// GetConsensusDemandRequest() 从对象池中获取ConsensusDemandRequest
func GetConsensusDemandRequest() *ConsensusDemandRequest {
	return poolConsensusDemandRequest.Get().(*ConsensusDemandRequest)
}

// ReleaseConsensusDemandRequest 释放ConsensusDemandRequest
func ReleaseConsensusDemandRequest(v *ConsensusDemandRequest) {
	v.ConsensusDemandParamDTOList = v.ConsensusDemandParamDTOList[:0]
	v.RequestExtendJson = ""
	poolConsensusDemandRequest.Put(v)
}
