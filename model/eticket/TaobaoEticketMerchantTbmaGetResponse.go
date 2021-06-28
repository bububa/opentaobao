package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码商查询淘宝码接口 APIResponse
taobao.eticket.merchant.tbma.get

码商查询淘宝码接口
*/
type TaobaoEticketMerchantTbmaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"eticket_merchant_tbma_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // respBody
    
    RespBody   *QueryTbMaCallbackResp `json:"resp_body,omitempty" xml:"