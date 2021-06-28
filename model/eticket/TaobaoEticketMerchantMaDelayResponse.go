package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
凭证延期 APIResponse
taobao.eticket.merchant.ma.delay

订单延期
*/
type TaobaoEticketMerchantMaDelayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"eticket_merchant_ma_delay_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"