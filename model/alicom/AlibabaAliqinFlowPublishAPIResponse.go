package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowPublishAPIResponse 流量发放(用户id) API返回值
// alibaba.aliqin.flow.publish
//
// 阿里通信流量下发功能，允许用户补发
type AlibabaAliqinFlowPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowPublishAPIResponseModel).Reset()
}

// AlibabaAliqinFlowPublishAPIResponseModel is 流量发放(用户id) 成功返回结果
type AlibabaAliqinFlowPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true为成功，其他为失败
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
}

var poolAlibabaAliqinFlowPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowPublishAPIResponse)
	},
}

// GetAlibabaAliqinFlowPublishAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowPublishAPIResponse
func GetAlibabaAliqinFlowPublishAPIResponse() *AlibabaAliqinFlowPublishAPIResponse {
	return poolAlibabaAliqinFlowPublishAPIResponse.Get().(*AlibabaAliqinFlowPublishAPIResponse)
}

// ReleaseAlibabaAliqinFlowPublishAPIResponse 将 AlibabaAliqinFlowPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowPublishAPIResponse(v *AlibabaAliqinFlowPublishAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowPublishAPIResponse.Put(v)
}
