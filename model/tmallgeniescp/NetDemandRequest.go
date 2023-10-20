package tmallgeniescp

import (
	"sync"
)

// NetDemandRequest 结构体
type NetDemandRequest struct {
	// 对象
	NetDemandDTOs []NetDemandDto `json:"net_demand_d_t_os,omitempty" xml:"net_demand_d_t_os>net_demand_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolNetDemandRequest = sync.Pool{
	New: func() any {
		return new(NetDemandRequest)
	},
}

// GetNetDemandRequest() 从对象池中获取NetDemandRequest
func GetNetDemandRequest() *NetDemandRequest {
	return poolNetDemandRequest.Get().(*NetDemandRequest)
}

// ReleaseNetDemandRequest 释放NetDemandRequest
func ReleaseNetDemandRequest(v *NetDemandRequest) {
	v.NetDemandDTOs = v.NetDemandDTOs[:0]
	v.RequestExtendJson = ""
	poolNetDemandRequest.Put(v)
}
