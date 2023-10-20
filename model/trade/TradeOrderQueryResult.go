package trade

import (
	"sync"
)

// TradeOrderQueryResult 结构体
type TradeOrderQueryResult struct {
	// 订单查询结果
	TradeList []Tradeorders `json:"trade_list,omitempty" xml:"trade_list>tradeorders,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 页码数量
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 数据数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 查询是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTradeOrderQueryResult = sync.Pool{
	New: func() any {
		return new(TradeOrderQueryResult)
	},
}

// GetTradeOrderQueryResult() 从对象池中获取TradeOrderQueryResult
func GetTradeOrderQueryResult() *TradeOrderQueryResult {
	return poolTradeOrderQueryResult.Get().(*TradeOrderQueryResult)
}

// ReleaseTradeOrderQueryResult 释放TradeOrderQueryResult
func ReleaseTradeOrderQueryResult(v *TradeOrderQueryResult) {
	v.TradeList = v.TradeList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageSize = 0
	v.PageIndex = 0
	v.PageCount = 0
	v.TotalCount = 0
	v.Success = false
	poolTradeOrderQueryResult.Put(v)
}
