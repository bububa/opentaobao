package aiar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ailab AR图像检索 APIResponse
alibaba.ai.ar.service.detect

ailab AR图像检索
*/
type AlibabaAiArServiceDetectAPIResponse struct {
    model.CommonResponse
    AlibabaAiArServiceDetectResponse
}

type AlibabaAiArServiceDetectResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_ar_service_detect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Results   string `json:"results,omitempty" xml:"results,omitempty"`

    
}
