package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全图文识别同步检测接口 APIResponse
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口
*/
type AlibabaSecurityJaqOcrImageSyncDetectAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqOcrImageSyncDetectResponse `json:"alibaba_security_jaq_ocr_image_sync_detect_response,omitempty"` 
    AlibabaSecurityJaqOcrImageSyncDetectResponse
}

/* model for simplify = false
type AlibabaSecurityJaqOcrImageSyncDetectResponse struct {

    // 出参结构体
    
    Data  *struct {
        JaqOcrImageDetectResult  *JaqOcrImageDetectResult `json:"jaq_ocr_image_detect_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqOcrImageSyncDetectResponse struct {

    // 出参结构体
    Data   *JaqOcrImageDetectResult `json:"data,omitempty"`

}
