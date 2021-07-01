package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqUrlScanAPIResponse
恶意网址检测接口 API返回值
alibaba.security.jaq.url.scan

url扫描接口 */
type AlibabaSecurityJaqUrlScanAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqUrlScanAPIResponseModel
}

// AlibabaSecurityJaqUrlScanAPIResponseModel is 恶意网址检测接口 成功返回结果
type AlibabaSecurityJaqUrlScanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_url_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扫描结果
	Data *UrlScanResult `json:"data,omitempty" xml:"data,omitempty"`
}
