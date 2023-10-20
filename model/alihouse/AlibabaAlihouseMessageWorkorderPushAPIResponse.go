package alihouse

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihouseMessageWorkorderPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseMessageWorkorderPushAPIResponseModel).Reset()
}

// AlibabaAlihouseMessageWorkorderPushAPIResponseModel is 工单消息推送 成功返回结果
type AlibabaAlihouseMessageWorkorderPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_message_workorder_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseMessageWorkorderPushResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseMessageWorkorderPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseMessageWorkorderPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMessageWorkorderPushAPIResponse)
	},
}

// GetAlibabaAlihouseMessageWorkorderPushAPIResponse 从 sync.Pool 获取 AlibabaAlihouseMessageWorkorderPushAPIResponse
func GetAlibabaAlihouseMessageWorkorderPushAPIResponse() *AlibabaAlihouseMessageWorkorderPushAPIResponse {
	return poolAlibabaAlihouseMessageWorkorderPushAPIResponse.Get().(*AlibabaAlihouseMessageWorkorderPushAPIResponse)
}

// ReleaseAlibabaAlihouseMessageWorkorderPushAPIResponse 将 AlibabaAlihouseMessageWorkorderPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseMessageWorkorderPushAPIResponse(v *AlibabaAlihouseMessageWorkorderPushAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseMessageWorkorderPushAPIResponse.Put(v)
}
