package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoNoteListResult 结构体
type TaobaoAilabAicloudTopMemoNoteListResult struct {
	// 服务的实际返回结果
	Notes []Note `json:"notes,omitempty" xml:"notes>note,omitempty"`
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoNoteListResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoNoteListResult)
	},
}

// GetTaobaoAilabAicloudTopMemoNoteListResult() 从对象池中获取TaobaoAilabAicloudTopMemoNoteListResult
func GetTaobaoAilabAicloudTopMemoNoteListResult() *TaobaoAilabAicloudTopMemoNoteListResult {
	return poolTaobaoAilabAicloudTopMemoNoteListResult.Get().(*TaobaoAilabAicloudTopMemoNoteListResult)
}

// ReleaseTaobaoAilabAicloudTopMemoNoteListResult 释放TaobaoAilabAicloudTopMemoNoteListResult
func ReleaseTaobaoAilabAicloudTopMemoNoteListResult(v *TaobaoAilabAicloudTopMemoNoteListResult) {
	v.Notes = v.Notes[:0]
	v.Message = ""
	v.StatusCode = 0
	poolTaobaoAilabAicloudTopMemoNoteListResult.Put(v)
}
