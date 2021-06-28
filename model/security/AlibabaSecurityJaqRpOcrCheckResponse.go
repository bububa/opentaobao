package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 APIResponse
alibaba.security.jaq.rp.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrCheckAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpOcrCheckResponse `json:"alibaba_security_jaq_rp_ocr_check_response,omitempty"` 
    AlibabaSecurityJaqRpOcrCheckResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpOcrCheckResponse struct {

    // result
    
    Data  *struct {
        RpidCard  *RpidCard `json:"rpid_card,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpOcrCheckResponse struct {

    // result
    Data   *RpidCard `json:"data,omitempty"`

}
