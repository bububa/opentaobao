package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场馆接口pushVenue APIResponse
alibaba.damai.mev.open.pushvenue

开放接口推送场馆
*/
type AlibabaDamaiMevOpenPushvenueAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushvenueResponse
}

type AlibabaDamaiMevOpenPushvenueResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushvenue_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushvenueResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
