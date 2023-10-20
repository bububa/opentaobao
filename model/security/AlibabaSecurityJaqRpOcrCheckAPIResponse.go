package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpOcrCheckAPIResponse ocr同时实名校验 API返回值
// alibaba.security.jaq.rp.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabaSecurityJaqRpOcrCheckAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpOcrCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpOcrCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpOcrCheckAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpOcrCheckAPIResponseModel is ocr同时实名校验 成功返回结果
type AlibabaSecurityJaqRpOcrCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_ocr_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpidCard `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpOcrCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpOcrCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpOcrCheckAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpOcrCheckAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpOcrCheckAPIResponse
func GetAlibabaSecurityJaqRpOcrCheckAPIResponse() *AlibabaSecurityJaqRpOcrCheckAPIResponse {
	return poolAlibabaSecurityJaqRpOcrCheckAPIResponse.Get().(*AlibabaSecurityJaqRpOcrCheckAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpOcrCheckAPIResponse 将 AlibabaSecurityJaqRpOcrCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpOcrCheckAPIResponse(v *AlibabaSecurityJaqRpOcrCheckAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpOcrCheckAPIResponse.Put(v)
}
