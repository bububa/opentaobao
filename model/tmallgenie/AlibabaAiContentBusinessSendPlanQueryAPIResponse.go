package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinesssendplanqueryAPIResponse 内容商业化发放权益查询 API返回值
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
type AlibabaaicontentbusinesssendplanqueryAPIResponse struct {
	model.CommonResponse
	AlibabaaicontentbusinesssendplanqueryAPIResponseModel
}

// AlibabaaicontentbusinesssendplanqueryAPIResponseModel is 内容商业化发放权益查询 成功返回结果
type AlibabaaicontentbusinesssendplanqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_content_business_send_plan_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	RetValue []PurchaseSendPlanDto `json:"ret_value,omitempty" xml:"ret_value>purchase_send_plan_dto,omitempty"`
	// 错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
