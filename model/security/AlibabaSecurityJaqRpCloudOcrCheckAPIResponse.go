package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudOcrCheckAPIResponse ocr同时实名校验 API返回值
// alibaba.security.jaq.rp.cloud.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabaSecurityJaqRpCloudOcrCheckAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudOcrCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudOcrCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudOcrCheckAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudOcrCheckAPIResponseModel is ocr同时实名校验 成功返回结果
type AlibabaSecurityJaqRpCloudOcrCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_ocr_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpidCard `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudOcrCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpCloudOcrCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudOcrCheckAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudOcrCheckAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudOcrCheckAPIResponse
func GetAlibabaSecurityJaqRpCloudOcrCheckAPIResponse() *AlibabaSecurityJaqRpCloudOcrCheckAPIResponse {
	return poolAlibabaSecurityJaqRpCloudOcrCheckAPIResponse.Get().(*AlibabaSecurityJaqRpCloudOcrCheckAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudOcrCheckAPIResponse 将 AlibabaSecurityJaqRpCloudOcrCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudOcrCheckAPIResponse(v *AlibabaSecurityJaqRpCloudOcrCheckAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudOcrCheckAPIResponse.Put(v)
}
