package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票品接口pushItem APIResponse
alibaba.damai.mev.open.pushitem

开放接口 推送票品
*/
type AlibabaDamaiMevOpenPushitemAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushitemResponse
}

type AlibabaDamaiMevOpenPushitemResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushitemResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
