package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
MT定制接口 APIResponse
alibaba.seaking.translate

MT定制接口
*/
type AlibabaSeakingTranslateAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingTranslateResponse
}

type AlibabaSeakingTranslateResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_translate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 译文
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
