package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票品接口deleteItem APIResponse
alibaba.damai.mev.open.deleteitem

deleteItem
*/
type AlibabaDamaiMevOpenDeleteitemAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenDeleteitemResponse
}

type AlibabaDamaiMevOpenDeleteitemResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenDeleteitemResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
