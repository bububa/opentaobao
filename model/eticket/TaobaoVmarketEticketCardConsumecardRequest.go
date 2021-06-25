package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证储值卡核销 APIRequest
taobao.vmarket.eticket.card.consumecard

线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
*/
type TaobaoVmarketEticketCardConsumecardRequest struct {
    model.Params

    // 卡号
    cardId   string 

    // 卡内等级
    cardLevel   int64 

    // 核销金额，精确到分，例如1.99元=199
    consumeValue   int64 

    // 买家昵称
    buyerNick   string 

    // 核销原因
    reason   string 

    // 门店id
    storeId   int64 

    // 操作人id
    operatorId   int64 

    // 核销流水号，外部ISV全局唯一
    consumeSerialNum   string 

    // 核销code
    consumeCode   string 

    // 安全token
    token   string 

}

func NewTaobaoVmarketEticketCardConsumecardRequest() *TaobaoVmarketEticketCardConsumecardRequest{
    return &TaobaoVmarketEticketCardConsumecardRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.card.consumecard"
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketCardConsumecardRequest) SetCardId(cardId string) error {
    r.cardId = cardId
    r.Set("card_id", cardId)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetCardId() string {
    return r.cardId
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetCardLevel(cardLevel int64) error {
    r.cardLevel = cardLevel
    r.Set("card_level", cardLevel)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetCardLevel() int64 {
    return r.cardLevel
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeValue(consumeValue int64) error {
    r.consumeValue = consumeValue
    r.Set("consume_value", consumeValue)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeValue() int64 {
    return r.consumeValue
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetOperatorId(operatorId int64) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetOperatorId() int64 {
    return r.operatorId
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeSerialNum(consumeSerialNum string) error {
    r.consumeSerialNum = consumeSerialNum
    r.Set("consume_serial_num", consumeSerialNum)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeSerialNum() string {
    return r.consumeSerialNum
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeCode(consumeCode string) error {
    r.consumeCode = consumeCode
    r.Set("consume_code", consumeCode)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeCode() string {
    return r.consumeCode
}

func (r *TaobaoVmarketEticketCardConsumecardRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoVmarketEticketCardConsumecardRequest) GetToken() string {
    return r.token
}

