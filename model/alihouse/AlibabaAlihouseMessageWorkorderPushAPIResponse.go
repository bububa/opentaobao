package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemessageworkorderpushAPIResponse 工单消息推送 API返回值
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
type AlibabaalihousemessageworkorderpushAPIResponse struct {
	model.CommonResponse
	AlibabaalihousemessageworkorderpushAPIResponseModel
}

// AlibabaalihousemessageworkorderpushAPIResponseModel is 工单消息推送 成功返回结果
type AlibabaalihousemessageworkorderpushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_message_workorder_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaalihousemessageworkorderpushResult `json:"result,omitempty" xml:"result,omitempty"`
}
