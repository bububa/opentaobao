package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppOfficialApplyAPIResponse 聚安全官方应用申请 API返回值
// alibaba.security.jaq.app.official.apply
//
// 官方应用申请接口
type AlibabaSecurityJaqAppOfficialApplyAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppOfficialApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppOfficialApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppOfficialApplyAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppOfficialApplyAPIResponseModel is 聚安全官方应用申请 成功返回结果
type AlibabaSecurityJaqAppOfficialApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_official_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 申请结果
	Result *OfficialAppApplyResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppOfficialApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppOfficialApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppOfficialApplyAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppOfficialApplyAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppOfficialApplyAPIResponse
func GetAlibabaSecurityJaqAppOfficialApplyAPIResponse() *AlibabaSecurityJaqAppOfficialApplyAPIResponse {
	return poolAlibabaSecurityJaqAppOfficialApplyAPIResponse.Get().(*AlibabaSecurityJaqAppOfficialApplyAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppOfficialApplyAPIResponse 将 AlibabaSecurityJaqAppOfficialApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppOfficialApplyAPIResponse(v *AlibabaSecurityJaqAppOfficialApplyAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppOfficialApplyAPIResponse.Put(v)
}
