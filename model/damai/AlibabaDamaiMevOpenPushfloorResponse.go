package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-楼层接口pushFloor APIResponse
alibaba.damai.mev.open.pushfloor

pushFloor
*/
type AlibabaDamaiMevOpenPushfloorAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushfloorResponse
}

type AlibabaDamaiMevOpenPushfloorResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushfloor_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushfloorResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
