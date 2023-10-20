package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqncopilotpictureauditAPIResponse AIGC创作图片审核 API返回值
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
type TaobaoqncopilotpictureauditAPIResponse struct {
	model.CommonResponse
	TaobaoqncopilotpictureauditAPIResponseModel
}

// TaobaoqncopilotpictureauditAPIResponseModel is AIGC创作图片审核 成功返回结果
type TaobaoqncopilotpictureauditAPIResponseModel struct {
	XMLName xml.Name `xml:"qncopilot_picture_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 审核结果
	QnCopilotResultDo *QnCopilotResultDo `json:"qn_copilot_result_do,omitempty" xml:"qn_copilot_result_do,omitempty"`
}
