package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全图文识别同步检测接口 APIResponse
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口
*/
type AlibabaSecurityJaqOcrImageSyncDetectAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqOcrImageSyncDetectResponse
}

type AlibabaSecurityJaqOcrImageSyncDetectResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_ocr_image_sync_detect_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参结构体
    
    Data   *JaqOcrImageDetectResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
