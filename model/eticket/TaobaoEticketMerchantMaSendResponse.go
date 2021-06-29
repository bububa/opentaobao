package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码成功回调接口 APIResponse
taobao.eticket.merchant.ma.send

码商发码成功回调接口
*/
type TaobaoEticketMerchantMaSendAPIResponse struct {
    model.CommonResponse
    TaobaoEticketMerchantMaSendResponse
}

type TaobaoEticketMerchantMaSendResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回复参数
    
    RespBody   *SendMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`

    
    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
