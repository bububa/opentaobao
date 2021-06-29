package aiar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ailab AR图像检索 API返回值 
alibaba.ai.ar.service.detect

ailab AR图像检索
*/
type AlibabaAiArServiceDetectAPIResponse struct {
    model.CommonResponse
    AlibabaAiArServiceDetectResponse
}

// ailab AR图像检索 成功返回结果
type AlibabaAiArServiceDetectResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_ar_service_detect_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Results   string `json:"results,omitempty" xml:"results,omitempty"`
}
