package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全智能鉴黄同步检测接口 APIResponse
alibaba.security.jaq.porn.image.sync.detect

同步黄图图像检测接口
*/
type AlibabaSecurityJaqPornImageSyncDetectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqPornImageSyncDetectResponse `json:"alibaba_security_jaq_porn_image_sync_detect_response,omitempty"`
}

type AlibabaSecurityJaqPornImageSyncDetectResponse struct {

    // 出参结构体
    Data   *JaqPornImageDetectResult `json:"data,omitempty"`

}
