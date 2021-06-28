package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的核销接口 APIResponse
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
*/
type TaobaoVmarketEticketAuthConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_auth_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:可以进行核销码操作
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"