package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子票券消费通知 APIResponse
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口
*/
type TaobaoVmarketEticketConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0:失败，1:成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"