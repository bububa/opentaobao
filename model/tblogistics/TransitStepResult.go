package tblogistics

import (
	"sync"
)

// TransitStepResult 结构体
type TransitStepResult struct {
	// 列表
	TraceList []TransitStepInfo `json:"trace_list,omitempty" xml:"trace_list>transit_step_info,omitempty"`
	// 运单号
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 订单物流状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolTransitStepResult = sync.Pool{
	New: func() any {
		return new(TransitStepResult)
	},
}

// GetTransitStepResult() 从对象池中获取TransitStepResult
func GetTransitStepResult() *TransitStepResult {
	return poolTransitStepResult.Get().(*TransitStepResult)
}

// ReleaseTransitStepResult 释放TransitStepResult
func ReleaseTransitStepResult(v *TransitStepResult) {
	v.TraceList = v.TraceList[:0]
	v.OutSid = ""
	v.CompanyName = ""
	v.Status = ""
	v.Tid = 0
	poolTransitStepResult.Put(v)
}
