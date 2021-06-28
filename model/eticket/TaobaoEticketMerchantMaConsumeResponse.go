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
	RequestId     string         `json:"request_id,omitempty" xml:"eticket_merchant_ma_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    RespBody   *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"