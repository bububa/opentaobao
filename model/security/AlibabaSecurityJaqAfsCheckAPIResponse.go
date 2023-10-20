package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAfsCheckAPIResponse 反欺诈二次验证接口 API返回值
// alibaba.security.jaq.afs.check
//
// 反欺诈二次验证接口
type AlibabaSecurityJaqAfsCheckAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAfsCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAfsCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAfsCheckAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAfsCheckAPIResponseModel is 反欺诈二次验证接口 成功返回结果
type AlibabaSecurityJaqAfsCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_afs_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验证结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAfsCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolAlibabaSecurityJaqAfsCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAfsCheckAPIResponse)
	},
}

// GetAlibabaSecurityJaqAfsCheckAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAfsCheckAPIResponse
func GetAlibabaSecurityJaqAfsCheckAPIResponse() *AlibabaSecurityJaqAfsCheckAPIResponse {
	return poolAlibabaSecurityJaqAfsCheckAPIResponse.Get().(*AlibabaSecurityJaqAfsCheckAPIResponse)
}

// ReleaseAlibabaSecurityJaqAfsCheckAPIResponse 将 AlibabaSecurityJaqAfsCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAfsCheckAPIResponse(v *AlibabaSecurityJaqAfsCheckAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAfsCheckAPIResponse.Put(v)
}
