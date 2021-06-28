package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证证件OCR识别功能接口 APIResponse
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpOcrResponse `json:"alibaba_security_jaq_rp_ocr_response,omitempty"` 
    AlibabaSecurityJaqRpOcrResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpOcrResponse struct {

    // 结果信息
    
    Data  *struct {
        RpidCardBo  *RpidCardBo `json:"rpid_card_bo,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpOcrResponse struct {

    // 结果信息
    Data   *RpidCardBo `json:"data,omitempty"`

}
