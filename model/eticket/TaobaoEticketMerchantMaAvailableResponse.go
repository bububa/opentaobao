package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销前校验接口 APIResponse
taobao.eticket.merchant.ma.available

商家验码之前的调用接口，用来判断是否可以进行核销操作
*/
type TaobaoEticketMerchantMaAvailableAPIResponse struct {
    model.CommonResponse
    TaobaoEticketMerchantMaAvailableResponse
}

type TaobaoEticketMerchantMaAvailableResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_available_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 回复结果
    
    RespBody   *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`

    
    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
