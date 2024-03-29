package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse 验证姓名和证件号 API返回值
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
type AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudRealnameCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudRealnameCheckAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudRealnameCheckAPIResponseModel is 验证姓名和证件号 成功返回结果
type AlibabaSecurityJaqRpCloudRealnameCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_realname_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RealNameResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudRealnameCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse
func GetAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse() *AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse {
	return poolAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse.Get().(*AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse 将 AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse(v *AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudRealnameCheckAPIResponse.Put(v)
}
