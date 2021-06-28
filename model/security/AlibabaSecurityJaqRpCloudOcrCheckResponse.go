package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 APIResponse
alibaba.security.jaq.rp.cloud.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpCloudOcrCheckAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudOcrCheckResponse `json:"alibaba_security_jaq_rp_cloud_ocr_check_response,omitempty"` 
    AlibabaSecurityJaqRpCloudOcrCheckResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudOcrCheckResponse struct {

    // result
    
    Data  *struct {
        RpidCard  *RpidCard `json:"rpid_card,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudOcrCheckResponse struct {

    // result
    Data   *RpidCard `json:"data,omitempty"`

}
