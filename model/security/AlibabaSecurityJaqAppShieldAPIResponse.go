package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppShieldAPIResponse 应用加固接口 API返回值
// alibaba.security.jaq.app.shield
//
// 提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
type AlibabaSecurityJaqAppShieldAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppShieldAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppShieldAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppShieldAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppShieldAPIResponseModel is 应用加固接口 成功返回结果
type AlibabaSecurityJaqAppShieldAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_shield_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加固任务信息
	Result *ScanTaskInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppShieldAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppShieldAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppShieldAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppShieldAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppShieldAPIResponse
func GetAlibabaSecurityJaqAppShieldAPIResponse() *AlibabaSecurityJaqAppShieldAPIResponse {
	return poolAlibabaSecurityJaqAppShieldAPIResponse.Get().(*AlibabaSecurityJaqAppShieldAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppShieldAPIResponse 将 AlibabaSecurityJaqAppShieldAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppShieldAPIResponse(v *AlibabaSecurityJaqAppShieldAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppShieldAPIResponse.Put(v)
}
