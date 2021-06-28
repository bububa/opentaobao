package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码失败回调接口 APIResponse
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoEticketMerchantMaFailsendAPIResponse struct {
    model.CommonResponse
    TaobaoEticketMerchantMaFailsendResponse
}

type TaobaoEticketMerchantMaFailsendResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_failsend_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 回复参数
    
    RespBody   *SendFailCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`

    
    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
