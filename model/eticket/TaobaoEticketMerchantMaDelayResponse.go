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
    TaobaoEticketMerchantMaDelayResponse
}

type TaobaoEticketMerchantMaDelayResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_delay_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    ResCode   int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`

    
    // 错误消息
    
    ResMsg   string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`

    
}
