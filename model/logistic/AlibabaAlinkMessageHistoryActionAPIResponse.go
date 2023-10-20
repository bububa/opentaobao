package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageHistoryActionAPIResponse 操作历史消息 API返回值
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
type AlibabaAlinkMessageHistoryActionAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkMessageHistoryActionAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryActionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkMessageHistoryActionAPIResponseModel).Reset()
}

// AlibabaAlinkMessageHistoryActionAPIResponseModel is 操作历史消息 成功返回结果
type AlibabaAlinkMessageHistoryActionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_history_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryActionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkMessageHistoryActionAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkMessageHistoryActionAPIResponse)
	},
}

// GetAlibabaAlinkMessageHistoryActionAPIResponse 从 sync.Pool 获取 AlibabaAlinkMessageHistoryActionAPIResponse
func GetAlibabaAlinkMessageHistoryActionAPIResponse() *AlibabaAlinkMessageHistoryActionAPIResponse {
	return poolAlibabaAlinkMessageHistoryActionAPIResponse.Get().(*AlibabaAlinkMessageHistoryActionAPIResponse)
}

// ReleaseAlibabaAlinkMessageHistoryActionAPIResponse 将 AlibabaAlinkMessageHistoryActionAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkMessageHistoryActionAPIResponse(v *AlibabaAlinkMessageHistoryActionAPIResponse) {
	v.Reset()
	poolAlibabaAlinkMessageHistoryActionAPIResponse.Put(v)
}
