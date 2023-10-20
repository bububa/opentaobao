package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoMeetingListAPIResponse 天猫精灵会议查询 API返回值
// taobao.ailab.aicloud.top.memo.meeting.list
//
// 查询天猫精灵用户设置的所有会议
type TaobaoAilabAicloudTopMemoMeetingListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoMeetingListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoMeetingListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMemoMeetingListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMemoMeetingListAPIResponseModel is 天猫精灵会议查询 成功返回结果
type TaobaoAilabAicloudTopMemoMeetingListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_meeting_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoAilabAicloudTopMemoMeetingListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoMeetingListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMemoMeetingListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoMeetingListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMemoMeetingListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoMeetingListAPIResponse
func GetTaobaoAilabAicloudTopMemoMeetingListAPIResponse() *TaobaoAilabAicloudTopMemoMeetingListAPIResponse {
	return poolTaobaoAilabAicloudTopMemoMeetingListAPIResponse.Get().(*TaobaoAilabAicloudTopMemoMeetingListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMemoMeetingListAPIResponse 将 TaobaoAilabAicloudTopMemoMeetingListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoMeetingListAPIResponse(v *TaobaoAilabAicloudTopMemoMeetingListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoMeetingListAPIResponse.Put(v)
}
