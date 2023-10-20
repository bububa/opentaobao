package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoAlarmListResult 结构体
type TaobaoAilabAicloudTopMemoAlarmListResult struct {
	// 服务的实际返回结果
	Alarms []Alarm `json:"alarms,omitempty" xml:"alarms>alarm,omitempty"`
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoAlarmListResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoAlarmListResult)
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmListResult() 从对象池中获取TaobaoAilabAicloudTopMemoAlarmListResult
func GetTaobaoAilabAicloudTopMemoAlarmListResult() *TaobaoAilabAicloudTopMemoAlarmListResult {
	return poolTaobaoAilabAicloudTopMemoAlarmListResult.Get().(*TaobaoAilabAicloudTopMemoAlarmListResult)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmListResult 释放TaobaoAilabAicloudTopMemoAlarmListResult
func ReleaseTaobaoAilabAicloudTopMemoAlarmListResult(v *TaobaoAilabAicloudTopMemoAlarmListResult) {
	v.Alarms = v.Alarms[:0]
	v.Message = ""
	v.StatusCode = 0
	poolTaobaoAilabAicloudTopMemoAlarmListResult.Put(v)
}
