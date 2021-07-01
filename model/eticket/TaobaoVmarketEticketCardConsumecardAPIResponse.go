package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证储值卡核销 API返回值 
taobao.vmarket.eticket.card.consumecard

线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
*/
type TaobaoVmarketEticketCardConsumecardAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketCardConsumecardAPIResponseModel
}

// 电子凭证储值卡核销 成功返回结果
type TaobaoVmarketEticketCardConsumecardAPIResponseModel struct {
    XMLName xml.Name `xml:"vmarket_eticket_card_consumecard_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 正确返回值
    Resultcode   int64 `json:"resultcode,omitempty" xml:"resultcode,omitempty"`
}
