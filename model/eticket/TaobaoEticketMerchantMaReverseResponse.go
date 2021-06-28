package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 APIResponse
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
type TaobaoEticketMerchantMaReverseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"eticket_merchant_ma_reverse_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 回复结果
    
    RespBody   *ReverseMaCallbackResp `json:"resp_body,omitempty" xml:"