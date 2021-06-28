package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证验码前置确认 APIResponse
taobao.vmarket.eticket.beforeconsume

商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
*/
type TaobaoVmarketEticketBeforeconsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_beforeconsume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:可以进行核销码操作
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"