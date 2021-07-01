package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse
天猫精灵会议删除 API返回值
taobao.ailab.aicloud.top.memo.meeting.delete

天猫精灵会议删除 */
type TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponseModel
}

// TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponseModel is 天猫精灵会议删除 成功返回结果
type TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_meeting_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoAilabAicloudTopMemoMeetingDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
