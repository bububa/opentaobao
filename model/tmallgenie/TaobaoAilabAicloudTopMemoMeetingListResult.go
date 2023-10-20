package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoMeetingListResult 结构体
type TaobaoAilabAicloudTopMemoMeetingListResult struct {
	// 服务的实际返回结果
	Meetings []Meeting `json:"meetings,omitempty" xml:"meetings>meeting,omitempty"`
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoMeetingListResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoMeetingListResult)
	},
}

// GetTaobaoAilabAicloudTopMemoMeetingListResult() 从对象池中获取TaobaoAilabAicloudTopMemoMeetingListResult
func GetTaobaoAilabAicloudTopMemoMeetingListResult() *TaobaoAilabAicloudTopMemoMeetingListResult {
	return poolTaobaoAilabAicloudTopMemoMeetingListResult.Get().(*TaobaoAilabAicloudTopMemoMeetingListResult)
}

// ReleaseTaobaoAilabAicloudTopMemoMeetingListResult 释放TaobaoAilabAicloudTopMemoMeetingListResult
func ReleaseTaobaoAilabAicloudTopMemoMeetingListResult(v *TaobaoAilabAicloudTopMemoMeetingListResult) {
	v.Meetings = v.Meetings[:0]
	v.Message = ""
	v.StatusCode = 0
	poolTaobaoAilabAicloudTopMemoMeetingListResult.Put(v)
}
