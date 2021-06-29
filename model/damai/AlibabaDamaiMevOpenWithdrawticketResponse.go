package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口withdrawTicket APIResponse
alibaba.damai.mev.open.withdrawticket

开放接口退票
*/
type AlibabaDamaiMevOpenWithdrawticketAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenWithdrawticketResponse
}

type AlibabaDamaiMevOpenWithdrawticketResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_withdrawticket_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenWithdrawticketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
