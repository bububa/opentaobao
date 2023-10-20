package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpOcrAPIResponse 聚安全实人认证证件OCR识别功能接口 API返回值
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabaSecurityJaqRpOcrAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpOcrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpOcrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpOcrAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpOcrAPIResponseModel is 聚安全实人认证证件OCR识别功能接口 成功返回结果
type AlibabaSecurityJaqRpOcrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_ocr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data *RpidCardBo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpOcrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpOcrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpOcrAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpOcrAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpOcrAPIResponse
func GetAlibabaSecurityJaqRpOcrAPIResponse() *AlibabaSecurityJaqRpOcrAPIResponse {
	return poolAlibabaSecurityJaqRpOcrAPIResponse.Get().(*AlibabaSecurityJaqRpOcrAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpOcrAPIResponse 将 AlibabaSecurityJaqRpOcrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpOcrAPIResponse(v *AlibabaSecurityJaqRpOcrAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpOcrAPIResponse.Put(v)
}
