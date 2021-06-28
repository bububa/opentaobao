package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务重新触发发码短信 APIResponse
taobao.vmarket.eticket.flow.resend

业务重新触发发码短信
*/
type TaobaoVmarketEticketFlowResendAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketFlowResendResponse
}

type TaobaoVmarketEticketFlowResendResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_flow_resend_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1成功;0失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 错误提示信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
