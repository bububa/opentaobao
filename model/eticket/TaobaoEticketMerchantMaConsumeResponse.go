package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销接口 APIResponse
taobao.eticket.merchant.ma.consume

电子凭证核销接口
*/
type TaobaoEticketMerchantMaConsumeAPIResponse struct {
    model.CommonResponse
    TaobaoEticketMerchantMaConsumeResponse
}

type TaobaoEticketMerchantMaConsumeResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_consume_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    RespBody   *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`

    
    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
