package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面接口pushFace APIResponse
alibaba.damai.mev.open.pushface

pushFace
*/
type AlibabaDamaiMevOpenPushfaceAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushfaceResponse
}

type AlibabaDamaiMevOpenPushfaceResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushface_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushfaceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
