package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageHistoryCountAPIResponse 查询消息总数 API返回值
// alibaba.alink.message.history.count
//
// 查询消息总数
type AlibabaAlinkMessageHistoryCountAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkMessageHistoryCountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkMessageHistoryCountAPIResponseModel).Reset()
}

// AlibabaAlinkMessageHistoryCountAPIResponseModel is 查询消息总数 成功返回结果
type AlibabaAlinkMessageHistoryCountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_history_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageHistoryCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkMessageHistoryCountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkMessageHistoryCountAPIResponse)
	},
}

// GetAlibabaAlinkMessageHistoryCountAPIResponse 从 sync.Pool 获取 AlibabaAlinkMessageHistoryCountAPIResponse
func GetAlibabaAlinkMessageHistoryCountAPIResponse() *AlibabaAlinkMessageHistoryCountAPIResponse {
	return poolAlibabaAlinkMessageHistoryCountAPIResponse.Get().(*AlibabaAlinkMessageHistoryCountAPIResponse)
}

// ReleaseAlibabaAlinkMessageHistoryCountAPIResponse 将 AlibabaAlinkMessageHistoryCountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkMessageHistoryCountAPIResponse(v *AlibabaAlinkMessageHistoryCountAPIResponse) {
	v.Reset()
	poolAlibabaAlinkMessageHistoryCountAPIResponse.Put(v)
}
