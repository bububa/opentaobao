package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
恶意网址检测接口 APIResponse
alibaba.security.jaq.url.scan

url扫描接口
*/
type AlibabaSecurityJaqUrlScanAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqUrlScanResponse `json:"alibaba_security_jaq_url_scan_response,omitempty"`
}

type AlibabaSecurityJaqUrlScanResponse struct {

    // 扫描结果
    Data   *UrlScanResult `json:"data,omitempty"`

}
