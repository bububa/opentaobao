package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudEventAPIResponse 事件上报 API返回值
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
type AlibabaSecurityJaqRpCloudEventAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudEventAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudEventAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudEventAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudEventAPIResponseModel is 事件上报 成功返回结果
type AlibabaSecurityJaqRpCloudEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RpEventResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudEventAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqRpCloudEventAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudEventAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudEventAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudEventAPIResponse
func GetAlibabaSecurityJaqRpCloudEventAPIResponse() *AlibabaSecurityJaqRpCloudEventAPIResponse {
	return poolAlibabaSecurityJaqRpCloudEventAPIResponse.Get().(*AlibabaSecurityJaqRpCloudEventAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudEventAPIResponse 将 AlibabaSecurityJaqRpCloudEventAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudEventAPIResponse(v *AlibabaSecurityJaqRpCloudEventAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudEventAPIResponse.Put(v)
}
