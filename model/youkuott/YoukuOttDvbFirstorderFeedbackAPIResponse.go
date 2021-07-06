package youkuott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttDvbFirstorderFeedbackAPIResponse dvb首次安装订单反馈 API返回值
// youku.ott.dvb.firstorder.feedback
//
// dvb首次安装订单反馈
type YoukuOttDvbFirstorderFeedbackAPIResponse struct {
	model.CommonResponse
	YoukuOttDvbFirstorderFeedbackAPIResponseModel
}

// YoukuOttDvbFirstorderFeedbackAPIResponseModel is dvb首次安装订单反馈 成功返回结果
type YoukuOttDvbFirstorderFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_dvb_firstorder_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功 true:成功 false:失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
