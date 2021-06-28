package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证证件OCR识别功能接口 APIResponse
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_ocr_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Data   *RpidCardBo `json:"data,omitempty" xml:"