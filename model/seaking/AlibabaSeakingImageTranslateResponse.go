package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片机器翻译 APIResponse
alibaba.seaking.imagetranslate

图片机器翻译
*/
type AlibabaSeakingImagetranslateAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingImagetranslateResponse
}

type AlibabaSeakingImagetranslateResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 译图url
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
