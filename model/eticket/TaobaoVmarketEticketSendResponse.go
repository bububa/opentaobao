package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家电子凭证发码成功回调接口 API返回值 
taobao.vmarket.eticket.send

外部商家成功发码回调接口
*/
type TaobaoVmarketEticketSendAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketSendResponse
}

// 商家电子凭证发码成功回调接口 成功返回结果
type TaobaoVmarketEticketSendResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 0:失败；1:成功
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
