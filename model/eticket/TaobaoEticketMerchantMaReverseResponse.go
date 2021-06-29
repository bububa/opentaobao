package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 API返回值 
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
type TaobaoEticketMerchantMaReverseAPIResponse struct {
    model.CommonResponse
    TaobaoEticketMerchantMaReverseResponse
}

// 电子凭证冲正接口 成功返回结果
type TaobaoEticketMerchantMaReverseResponse struct {
    XMLName xml.Name `xml:"eticket_merchant_ma_reverse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回复结果
    RespBody   *ReverseMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
    // 子结果码
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
}
