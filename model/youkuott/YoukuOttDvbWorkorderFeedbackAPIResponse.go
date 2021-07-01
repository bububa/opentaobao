package youkuott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttDvbWorkorderFeedbackAPIResponse
dvb工单反馈 API返回值
youku.ott.dvb.workorder.feedback

dvb工单处理结果反馈 */
type YoukuOttDvbWorkorderFeedbackAPIResponse struct {
	model.CommonResponse
	YoukuOttDvbWorkorderFeedbackAPIResponseModel
}

// YoukuOttDvbWorkorderFeedbackAPIResponseModel is dvb工单反馈 成功返回结果
type YoukuOttDvbWorkorderFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_dvb_workorder_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功 true:成功 false:失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
