package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse 天猫精灵闹钟创建 API返回值
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
type TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoAlarmCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMemoAlarmCreateAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMemoAlarmCreateAPIResponseModel is 天猫精灵闹钟创建 成功返回结果
type TaobaoAilabAicloudTopMemoAlarmCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 闹钟ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoAlarmCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Result = 0
	m.StatusCode = 0
}

var poolTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse
func GetTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse() *TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse {
	return poolTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse.Get().(*TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse 将 TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse(v *TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoAlarmCreateAPIResponse.Put(v)
}
