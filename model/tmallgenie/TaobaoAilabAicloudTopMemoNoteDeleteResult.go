package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoNoteDeleteResult 结构体
type TaobaoAilabAicloudTopMemoNoteDeleteResult struct {
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 服务的实际返回结果
	Note *Note `json:"note,omitempty" xml:"note,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoNoteDeleteResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoNoteDeleteResult)
	},
}

// GetTaobaoAilabAicloudTopMemoNoteDeleteResult() 从对象池中获取TaobaoAilabAicloudTopMemoNoteDeleteResult
func GetTaobaoAilabAicloudTopMemoNoteDeleteResult() *TaobaoAilabAicloudTopMemoNoteDeleteResult {
	return poolTaobaoAilabAicloudTopMemoNoteDeleteResult.Get().(*TaobaoAilabAicloudTopMemoNoteDeleteResult)
}

// ReleaseTaobaoAilabAicloudTopMemoNoteDeleteResult 释放TaobaoAilabAicloudTopMemoNoteDeleteResult
func ReleaseTaobaoAilabAicloudTopMemoNoteDeleteResult(v *TaobaoAilabAicloudTopMemoNoteDeleteResult) {
	v.Message = ""
	v.StatusCode = 0
	v.Note = nil
	poolTaobaoAilabAicloudTopMemoNoteDeleteResult.Put(v)
}
