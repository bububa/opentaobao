package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMessageWorkorderPushAPIResponse 工单消息推送 API返回值
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
type AlibabaAlihouseMessageWorkorderPushAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseMessageWorkorderPushAPIResponseModel
}

// AlibabaAlihouseMessageWorkorderPushAPIResponseModel is 工单消息推送 成功返回结果
type AlibabaAlihouseMessageWorkorderPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_message_workorder_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseMessageWorkorderPushResult `json:"result,omitempty" xml:"result,omitempty"`
}
