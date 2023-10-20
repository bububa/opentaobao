package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmListAPIResponse 天猫精灵闹钟查询 API返回值
// taobao.ailab.aicloud.top.memo.alarm.list
//
// 查询天猫精灵用户设置的所有闹钟
type TaobaoAilabAicloudTopMemoAlarmListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoAlarmListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoAlarmListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMemoAlarmListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMemoAlarmListAPIResponseModel is 天猫精灵闹钟查询 成功返回结果
type TaobaoAilabAicloudTopMemoAlarmListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoAilabAicloudTopMemoAlarmListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoAlarmListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMemoAlarmListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoAlarmListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoAlarmListAPIResponse
func GetTaobaoAilabAicloudTopMemoAlarmListAPIResponse() *TaobaoAilabAicloudTopMemoAlarmListAPIResponse {
	return poolTaobaoAilabAicloudTopMemoAlarmListAPIResponse.Get().(*TaobaoAilabAicloudTopMemoAlarmListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmListAPIResponse 将 TaobaoAilabAicloudTopMemoAlarmListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoAlarmListAPIResponse(v *TaobaoAilabAicloudTopMemoAlarmListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoAlarmListAPIResponse.Put(v)
}
