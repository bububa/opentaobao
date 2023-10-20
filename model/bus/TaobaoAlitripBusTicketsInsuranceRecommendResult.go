package bus

import (
	"sync"
)

// TaobaoAlitripBusTicketsInsuranceRecommendResult 结构体
type TaobaoAlitripBusTicketsInsuranceRecommendResult struct {
	// 扩展预留
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口返回结果对象
	Response *TopStandardInsRecommendResponse `json:"response,omitempty" xml:"response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripBusTicketsInsuranceRecommendResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripBusTicketsInsuranceRecommendResult)
	},
}

// GetTaobaoAlitripBusTicketsInsuranceRecommendResult() 从对象池中获取TaobaoAlitripBusTicketsInsuranceRecommendResult
func GetTaobaoAlitripBusTicketsInsuranceRecommendResult() *TaobaoAlitripBusTicketsInsuranceRecommendResult {
	return poolTaobaoAlitripBusTicketsInsuranceRecommendResult.Get().(*TaobaoAlitripBusTicketsInsuranceRecommendResult)
}

// ReleaseTaobaoAlitripBusTicketsInsuranceRecommendResult 释放TaobaoAlitripBusTicketsInsuranceRecommendResult
func ReleaseTaobaoAlitripBusTicketsInsuranceRecommendResult(v *TaobaoAlitripBusTicketsInsuranceRecommendResult) {
	v.BizExtMap = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Response = nil
	v.Success = false
	poolTaobaoAlitripBusTicketsInsuranceRecommendResult.Put(v)
}
