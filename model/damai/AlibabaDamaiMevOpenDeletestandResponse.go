package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口deleteStand APIResponse
alibaba.damai.mev.open.deletestand

deleteStand
*/
type AlibabaDamaiMevOpenDeletestandAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenDeletestandResponse
}

type AlibabaDamaiMevOpenDeletestandResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_deletestand_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenDeletestandResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
