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
    // Response *AlibabaSecurityJaqUrlScanResponse `json:"alibaba_security_jaq_url_scan_response,omitempty"` 
    AlibabaSecurityJaqUrlScanResponse
}

/* model for simplify = false
type AlibabaSecurityJaqUrlScanResponse struct {

    // 扫描结果
    
    Data  *struct {
        UrlScanResult  *UrlScanResult `json:"url_scan_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqUrlScanResponse struct {

    // 扫描结果
    Data   *UrlScanResult `json:"data,omitempty"`

}
