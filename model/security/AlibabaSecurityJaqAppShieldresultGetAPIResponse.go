package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppShieldresultGetAPIResponse 用户查询加固结果 API返回值
// alibaba.security.jaq.app.shieldresult.get
//
// 用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
type AlibabaSecurityJaqAppShieldresultGetAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppShieldresultGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppShieldresultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppShieldresultGetAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppShieldresultGetAPIResponseModel is 用户查询加固结果 成功返回结果
type AlibabaSecurityJaqAppShieldresultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_shieldresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应用加固结果
	Result *ShieldResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppShieldresultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppShieldresultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppShieldresultGetAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppShieldresultGetAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppShieldresultGetAPIResponse
func GetAlibabaSecurityJaqAppShieldresultGetAPIResponse() *AlibabaSecurityJaqAppShieldresultGetAPIResponse {
	return poolAlibabaSecurityJaqAppShieldresultGetAPIResponse.Get().(*AlibabaSecurityJaqAppShieldresultGetAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppShieldresultGetAPIResponse 将 AlibabaSecurityJaqAppShieldresultGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppShieldresultGetAPIResponse(v *AlibabaSecurityJaqAppShieldresultGetAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppShieldresultGetAPIResponse.Put(v)
}
