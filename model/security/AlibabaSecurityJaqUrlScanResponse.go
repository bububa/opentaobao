package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
恶意网址检测接口 APIResponse
alibaba.security.jaq.url.scan

url扫描接口
*/
type AlibabaSecurityJaqUrlScanAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqUrlScanResponse
}

type AlibabaSecurityJaqUrlScanResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_url_scan_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 扫描结果
    
    Data   *UrlScanResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
