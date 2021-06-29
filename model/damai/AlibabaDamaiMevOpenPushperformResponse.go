package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场次接口pushPerform APIResponse
alibaba.damai.mev.open.pushperform

pushPerform
*/
type AlibabaDamaiMevOpenPushperformAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushperformResponse
}

type AlibabaDamaiMevOpenPushperformResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushperform_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushperformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
