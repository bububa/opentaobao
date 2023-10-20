package tblogistics

import (
	"sync"
)

// TaobaoLogisticsInstantTraceSearchResult 结构体
type TaobaoLogisticsInstantTraceSearchResult struct {
	// 运单列表
	MailList []TopLogisticsMailDto `json:"mail_list,omitempty" xml:"mail_list>top_logistics_mail_dto,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoLogisticsInstantTraceSearchResult = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsInstantTraceSearchResult)
	},
}

// GetTaobaoLogisticsInstantTraceSearchResult() 从对象池中获取TaobaoLogisticsInstantTraceSearchResult
func GetTaobaoLogisticsInstantTraceSearchResult() *TaobaoLogisticsInstantTraceSearchResult {
	return poolTaobaoLogisticsInstantTraceSearchResult.Get().(*TaobaoLogisticsInstantTraceSearchResult)
}

// ReleaseTaobaoLogisticsInstantTraceSearchResult 释放TaobaoLogisticsInstantTraceSearchResult
func ReleaseTaobaoLogisticsInstantTraceSearchResult(v *TaobaoLogisticsInstantTraceSearchResult) {
	v.MailList = v.MailList[:0]
	v.Success = false
	poolTaobaoLogisticsInstantTraceSearchResult.Put(v)
}
