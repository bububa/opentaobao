package traveltrade

import (
	"sync"
)

// AlitripTravelTradeTemplateQueryResultSet 结构体
type AlitripTravelTradeTemplateQueryResultSet struct {
	// 订单服务标注模版获取结果详情
	FirstResult *OrderTipInfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

var poolAlitripTravelTradeTemplateQueryResultSet = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeTemplateQueryResultSet)
	},
}

// GetAlitripTravelTradeTemplateQueryResultSet() 从对象池中获取AlitripTravelTradeTemplateQueryResultSet
func GetAlitripTravelTradeTemplateQueryResultSet() *AlitripTravelTradeTemplateQueryResultSet {
	return poolAlitripTravelTradeTemplateQueryResultSet.Get().(*AlitripTravelTradeTemplateQueryResultSet)
}

// ReleaseAlitripTravelTradeTemplateQueryResultSet 释放AlitripTravelTradeTemplateQueryResultSet
func ReleaseAlitripTravelTradeTemplateQueryResultSet(v *AlitripTravelTradeTemplateQueryResultSet) {
	v.FirstResult = nil
	poolAlitripTravelTradeTemplateQueryResultSet.Put(v)
}
