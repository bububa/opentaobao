package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无法发码回调 APIResponse
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoVmarketEticketFailsendAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketFailsendResponse
}

type TaobaoVmarketEticketFailsendResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_failsend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
}
