package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppOfficialVerifyAPIResponse 聚安全验证官方应用接口 API返回值
// alibaba.security.jaq.app.official.verify
//
// 接入用户来查询应用是否为官方应用
type AlibabaSecurityJaqAppOfficialVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppOfficialVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppOfficialVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppOfficialVerifyAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppOfficialVerifyAPIResponseModel is 聚安全验证官方应用接口 成功返回结果
type AlibabaSecurityJaqAppOfficialVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_official_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OfficialAppVerifyResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppOfficialVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppOfficialVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppOfficialVerifyAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppOfficialVerifyAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppOfficialVerifyAPIResponse
func GetAlibabaSecurityJaqAppOfficialVerifyAPIResponse() *AlibabaSecurityJaqAppOfficialVerifyAPIResponse {
	return poolAlibabaSecurityJaqAppOfficialVerifyAPIResponse.Get().(*AlibabaSecurityJaqAppOfficialVerifyAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppOfficialVerifyAPIResponse 将 AlibabaSecurityJaqAppOfficialVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppOfficialVerifyAPIResponse(v *AlibabaSecurityJaqAppOfficialVerifyAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppOfficialVerifyAPIResponse.Put(v)
}
