package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopIdListConverterResult 结构体
type TaobaoAilabAicloudTopIdListConverterResult struct {
	// 返回查询内容
	RtValue []OpenInfoResponse `json:"rt_value,omitempty" xml:"rt_value>open_info_response,omitempty"`
	// 返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回code
	RtCode int64 `json:"rt_code,omitempty" xml:"rt_code,omitempty"`
	// 请求状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTaobaoAilabAicloudTopIdListConverterResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopIdListConverterResult)
	},
}

// GetTaobaoAilabAicloudTopIdListConverterResult() 从对象池中获取TaobaoAilabAicloudTopIdListConverterResult
func GetTaobaoAilabAicloudTopIdListConverterResult() *TaobaoAilabAicloudTopIdListConverterResult {
	return poolTaobaoAilabAicloudTopIdListConverterResult.Get().(*TaobaoAilabAicloudTopIdListConverterResult)
}

// ReleaseTaobaoAilabAicloudTopIdListConverterResult 释放TaobaoAilabAicloudTopIdListConverterResult
func ReleaseTaobaoAilabAicloudTopIdListConverterResult(v *TaobaoAilabAicloudTopIdListConverterResult) {
	v.RtValue = v.RtValue[:0]
	v.Message = ""
	v.RtCode = 0
	v.IsSuccess = false
	poolTaobaoAilabAicloudTopIdListConverterResult.Put(v)
}
