package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoAlarmDeleteResult 结构体
type TaobaoAilabAicloudTopMemoAlarmDeleteResult struct {
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 服务的实际返回结果
	Alarm *Alarm `json:"alarm,omitempty" xml:"alarm,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoAlarmDeleteResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoAlarmDeleteResult)
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmDeleteResult() 从对象池中获取TaobaoAilabAicloudTopMemoAlarmDeleteResult
func GetTaobaoAilabAicloudTopMemoAlarmDeleteResult() *TaobaoAilabAicloudTopMemoAlarmDeleteResult {
	return poolTaobaoAilabAicloudTopMemoAlarmDeleteResult.Get().(*TaobaoAilabAicloudTopMemoAlarmDeleteResult)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmDeleteResult 释放TaobaoAilabAicloudTopMemoAlarmDeleteResult
func ReleaseTaobaoAilabAicloudTopMemoAlarmDeleteResult(v *TaobaoAilabAicloudTopMemoAlarmDeleteResult) {
	v.Message = ""
	v.StatusCode = 0
	v.Alarm = nil
	poolTaobaoAilabAicloudTopMemoAlarmDeleteResult.Put(v)
}
