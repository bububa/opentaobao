package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqUrlScanAPIResponse 恶意网址检测接口 API返回值
// alibaba.security.jaq.url.scan
//
// url扫描接口
type AlibabaSecurityJaqUrlScanAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqUrlScanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqUrlScanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqUrlScanAPIResponseModel).Reset()
}

// AlibabaSecurityJaqUrlScanAPIResponseModel is 恶意网址检测接口 成功返回结果
type AlibabaSecurityJaqUrlScanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_url_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扫描结果
	Data *UrlScanResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqUrlScanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqUrlScanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqUrlScanAPIResponse)
	},
}

// GetAlibabaSecurityJaqUrlScanAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqUrlScanAPIResponse
func GetAlibabaSecurityJaqUrlScanAPIResponse() *AlibabaSecurityJaqUrlScanAPIResponse {
	return poolAlibabaSecurityJaqUrlScanAPIResponse.Get().(*AlibabaSecurityJaqUrlScanAPIResponse)
}

// ReleaseAlibabaSecurityJaqUrlScanAPIResponse 将 AlibabaSecurityJaqUrlScanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqUrlScanAPIResponse(v *AlibabaSecurityJaqUrlScanAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqUrlScanAPIResponse.Put(v)
}
