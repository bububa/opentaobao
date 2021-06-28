package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无交易类凭证核销 APIResponse
taobao.vmarket.eticket.flow.consume

无交易类凭证核销
*/
type TaobaoVmarketEticketFlowConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_flow_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 执行成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"