package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全智能鉴黄同步检测接口 APIResponse
alibaba.security.jaq.porn.image.sync.detect

同步黄图图像检测接口
*/
type AlibabaSecurityJaqPornImageSyncDetectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_porn_image_sync_detect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参结构体
    
    Data   *JaqPornImageDetectResult `json:"data,omitempty" xml:"