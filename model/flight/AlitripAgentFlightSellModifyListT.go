package flight

import (
	"sync"
)

// AlitripAgentFlightSellModifyListT 结构体
type AlitripAgentFlightSellModifyListT struct {
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 国内国际标识(1:国内,2:国际)
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolAlitripAgentFlightSellModifyListT = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyListT)
	},
}

// GetAlitripAgentFlightSellModifyListT() 从对象池中获取AlitripAgentFlightSellModifyListT
func GetAlitripAgentFlightSellModifyListT() *AlitripAgentFlightSellModifyListT {
	return poolAlitripAgentFlightSellModifyListT.Get().(*AlitripAgentFlightSellModifyListT)
}

// ReleaseAlitripAgentFlightSellModifyListT 释放AlitripAgentFlightSellModifyListT
func ReleaseAlitripAgentFlightSellModifyListT(v *AlitripAgentFlightSellModifyListT) {
	v.ApplyId = ""
	v.OrderId = ""
	v.ApplyTime = ""
	v.DomesticIntl = 0
	poolAlitripAgentFlightSellModifyListT.Put(v)
}
