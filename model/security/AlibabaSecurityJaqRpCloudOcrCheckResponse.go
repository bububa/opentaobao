package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 APIResponse
alibaba.security.jaq.rp.cloud.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpCloudOcrCheckAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpCloudOcrCheckResponse
}

type AlibabaSecurityJaqRpCloudOcrCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_ocr_check_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Data   *RpidCard `json:"data,omitempty" xml:"data,omitempty"`

    
}
