package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessSendPlanReceiveAPIResponse 天猫精灵商业化采销发放计划领取 API返回值
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
type AlibabaAiContentBusinessSendPlanReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaAiContentBusinessSendPlanReceiveAPIResponseModel
}

// AlibabaAiContentBusinessSendPlanReceiveAPIResponseModel is 天猫精灵商业化采销发放计划领取 成功返回结果
type AlibabaAiContentBusinessSendPlanReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_content_business_send_plan_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 返回对象
	RetValue *PurchaseReceiveReturnDto `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
