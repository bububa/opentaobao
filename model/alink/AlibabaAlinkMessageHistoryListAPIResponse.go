package alink

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkMessageHistoryListAPIResponseModel).Reset()
}

// AlibabaAlinkMessageHistoryListAPIResponseModel is 查询消息列表 成功返回结果
type AlibabaAlinkMessageHistoryListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_history_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkMessageHistoryListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkMessageHistoryListAPIResponse)
	},
}

// GetAlibabaAlinkMessageHistoryListAPIResponse 从 sync.Pool 获取 AlibabaAlinkMessageHistoryListAPIResponse
func GetAlibabaAlinkMessageHistoryListAPIResponse() *AlibabaAlinkMessageHistoryListAPIResponse {
	return poolAlibabaAlinkMessageHistoryListAPIResponse.Get().(*AlibabaAlinkMessageHistoryListAPIResponse)
}

// ReleaseAlibabaAlinkMessageHistoryListAPIResponse 将 AlibabaAlinkMessageHistoryListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkMessageHistoryListAPIResponse(v *AlibabaAlinkMessageHistoryListAPIResponse) {
	v.Reset()
	poolAlibabaAlinkMessageHistoryListAPIResponse.Put(v)
}
