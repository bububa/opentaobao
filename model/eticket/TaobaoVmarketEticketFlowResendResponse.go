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
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_flow_resend_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1成功;0失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"