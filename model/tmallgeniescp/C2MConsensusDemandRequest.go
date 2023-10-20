package tmallgeniescp

import (
	"sync"
)

// C2MConsensusDemandRequest 结构体
type C2MConsensusDemandRequest struct {
	// 对象
	C2MConsensusDemandParamDTOList []C2MConsensusDemandParamDto `json:"c2_m_consensus_demand_param_d_t_o_list,omitempty" xml:"c2_m_consensus_demand_param_d_t_o_list>c2m_consensus_demand_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolC2MConsensusDemandRequest = sync.Pool{
	New: func() any {
		return new(C2MConsensusDemandRequest)
	},
}

// GetC2MConsensusDemandRequest() 从对象池中获取C2MConsensusDemandRequest
func GetC2MConsensusDemandRequest() *C2MConsensusDemandRequest {
	return poolC2MConsensusDemandRequest.Get().(*C2MConsensusDemandRequest)
}

// ReleaseC2MConsensusDemandRequest 释放C2MConsensusDemandRequest
func ReleaseC2MConsensusDemandRequest(v *C2MConsensusDemandRequest) {
	v.C2MConsensusDemandParamDTOList = v.C2MConsensusDemandParamDTOList[:0]
	v.RequestExtendJson = ""
	poolC2MConsensusDemandRequest.Put(v)
}
