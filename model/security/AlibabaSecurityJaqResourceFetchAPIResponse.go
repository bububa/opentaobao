package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqResourceFetchAPIResponse 获取资源文件 API返回值
// alibaba.security.jaq.resource.fetch
//
// 在前向化验证流程中提供资源文件服务
type AlibabaSecurityJaqResourceFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqResourceFetchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqResourceFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqResourceFetchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqResourceFetchAPIResponseModel is 获取资源文件 成功返回结果
type AlibabaSecurityJaqResourceFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_resource_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取资源结果
	Data *JaqResourceResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqResourceFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqResourceFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqResourceFetchAPIResponse)
	},
}

// GetAlibabaSecurityJaqResourceFetchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqResourceFetchAPIResponse
func GetAlibabaSecurityJaqResourceFetchAPIResponse() *AlibabaSecurityJaqResourceFetchAPIResponse {
	return poolAlibabaSecurityJaqResourceFetchAPIResponse.Get().(*AlibabaSecurityJaqResourceFetchAPIResponse)
}

// ReleaseAlibabaSecurityJaqResourceFetchAPIResponse 将 AlibabaSecurityJaqResourceFetchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqResourceFetchAPIResponse(v *AlibabaSecurityJaqResourceFetchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqResourceFetchAPIResponse.Put(v)
}
