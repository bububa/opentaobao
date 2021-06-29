package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口lockTicket APIResponse
alibaba.damai.mev.open.lockticket

开放接口 冻结票单
*/
type AlibabaDamaiMevOpenLockticketAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenLockticketResponse
}

type AlibabaDamaiMevOpenLockticketResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_lockticket_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenLockticketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
