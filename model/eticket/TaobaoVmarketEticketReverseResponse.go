package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 APIResponse
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口
*/
type TaobaoVmarketEticketReverseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_reverse_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0:失败，1:成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"