package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageHistoryListAPIResponse 查询消息列表 API返回值
// alibaba.alink.message.history.list
//
// 查询消息列表
type AlibabaAlinkMessageHistoryListAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkMessageHistoryListAPIResponseModel
}

// AlibabaAlinkMessageHistoryListAPIResponseModel is 查询消息列表 成功返回结果
type AlibabaAlinkMessageHistoryListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_history_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
