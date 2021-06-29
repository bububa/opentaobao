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
    TaobaoEticketMerchantTbmaGetResponse
}

type TaobaoEticketMerchantTbmaGetResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_tbma_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // respBody
    
    RespBody   *QueryTbMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`

    
    // subCode
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // subMsg
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
