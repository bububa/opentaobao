package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkmessagehistoryactionAPIResponse 操作历史消息 API返回值
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
type AlibabaalinkmessagehistoryactionAPIResponse struct {
	model.CommonResponse
	AlibabaalinkmessagehistoryactionAPIResponseModel
}

// AlibabaalinkmessagehistoryactionAPIResponseModel is 操作历史消息 成功返回结果
type AlibabaalinkmessagehistoryactionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_history_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
