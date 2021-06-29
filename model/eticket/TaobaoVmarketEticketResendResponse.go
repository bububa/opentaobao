package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部合作商家重发电子凭证回调接口 APIResponse
taobao.vmarket.eticket.resend

外部合作商家重发电子凭证回调接口
*/
type TaobaoVmarketEticketResendAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketResendResponse
}

type TaobaoVmarketEticketResendResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_resend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 0:失败，1:成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
}
