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
	RequestId     string         `json:"request_id,omitempty" xml:"eticket_merchant_ma_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 回复参数
    
    RespBody   *SendMaCallbackResp `json:"resp_body,omitempty" xml:"