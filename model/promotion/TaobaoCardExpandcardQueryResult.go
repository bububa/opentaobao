package promotion

import (
	"sync"
)

// TaobaoCardExpandcardQueryResult 结构体
type TaobaoCardExpandcardQueryResult struct {
	// 卡信息
	Models []ExpandCardVo `json:"models,omitempty" xml:"models>expand_card_vo,omitempty"`
	// debugInfo
	DebugInfo string `json:"debug_info,omitempty" xml:"debug_info,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误级别
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 0为成功，其他为失败
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoCardExpandcardQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoCardExpandcardQueryResult)
	},
}

// GetTaobaoCardExpandcardQueryResult() 从对象池中获取TaobaoCardExpandcardQueryResult
func GetTaobaoCardExpandcardQueryResult() *TaobaoCardExpandcardQueryResult {
	return poolTaobaoCardExpandcardQueryResult.Get().(*TaobaoCardExpandcardQueryResult)
}

// ReleaseTaobaoCardExpandcardQueryResult 释放TaobaoCardExpandcardQueryResult
func ReleaseTaobaoCardExpandcardQueryResult(v *TaobaoCardExpandcardQueryResult) {
	v.Models = v.Models[:0]
	v.DebugInfo = ""
	v.Message = ""
	v.ErrorLevel = ""
	v.Code = 0
	poolTaobaoCardExpandcardQueryResult.Put(v)
}
