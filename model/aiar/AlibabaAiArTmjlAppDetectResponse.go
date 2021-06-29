package aiar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵扫一扫入口的服务 APIResponse
alibaba.ai.ar.tmjl.app.detect

天猫精灵扫一扫入口的图像检测服务
*/
type AlibabaAiArTmjlAppDetectAPIResponse struct {
    model.CommonResponse
    AlibabaAiArTmjlAppDetectResponse
}

type AlibabaAiArTmjlAppDetectResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_ar_tmjl_app_detect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Results   string `json:"results,omitempty" xml:"results,omitempty"`

    
}
