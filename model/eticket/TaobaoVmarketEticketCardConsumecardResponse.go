package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证储值卡核销 APIResponse
taobao.vmarket.eticket.card.consumecard

线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
*/
type TaobaoVmarketEticketCardConsumecardAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketCardConsumecardResponse `json:"taobao_vmarket_eticket_card_consumecard_response,omitempty"`
}

type TaobaoVmarketEticketCardConsumecardResponse struct {

    // 正确返回值
    Resultcode   int64 `json:"resultcode,omitempty"`

}
