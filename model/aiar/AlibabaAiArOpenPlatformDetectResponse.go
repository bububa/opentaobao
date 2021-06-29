package aiar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AR开发者平台marker图片检测服务 APIResponse
alibaba.ai.ar.open.platform.detect

AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用
*/
type AlibabaAiArOpenPlatformDetectAPIResponse struct {
    model.CommonResponse
    AlibabaAiArOpenPlatformDetectResponse
}

type AlibabaAiArOpenPlatformDetectResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_ar_open_platform_detect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
