package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 API返回值 
alibaba.security.jaq.rp.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrCheckAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpOcrCheckResponse
}

// ocr同时实名校验 成功返回结果
type AlibabaSecurityJaqRpOcrCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_ocr_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Data   *RpidCard `json:"data,omitempty" xml:"data,omitempty"`
}
