package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片语种识别 APIResponse
alibaba.seaking.imagerecognize

图片语种识别
*/
type AlibabaSeakingImagerecognizeAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingImagerecognizeResponse
}

type AlibabaSeakingImagerecognizeResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_imagerecognize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 识别出的图片语种
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
