package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证储值卡核销 API请求
taobao.vmarket.eticket.card.consumecard

线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
*/
type TaobaoVmarketEticketCardConsumecardRequest struct {
    model.Params
    // 卡号
    _cardId   string
    // 卡内等级
    _cardLevel   int64
    // 核销金额，精确到分，例如1.99元=199
    _consumeValue   int64
    // 买家昵称
    _buyerNick   string
    // 核销原因
    _reason   string
    // 门店id
    _storeId   int64
    // 操作人id
    _operatorId   int64
    // 核销流水号，外部ISV全局唯一
    _consumeSerialNum   string
    // 核销code
    _consumeCode   string
    // 安全token
    _token   string
}

// 初始化TaobaoVmarketEticketCardConsumecardRequest对象
func NewTaobaoVmarketEticketCardConsumecardRequest() *TaobaoVmarketEticketCardConsumecardRequest{
    return &TaobaoVmarketEticketCardConsumecardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketCardConsumecardRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.card.consumecard"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketCardConsumecardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CardId Setter
// 卡号
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetCardId(_cardId string) error {
    r._cardId = _cardId
    r.Set("card_id", _cardId)
    return nil
}

// CardId Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetCardId() string {
    return r._cardId
}
// CardLevel Setter
// 卡内等级
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetCardLevel(_cardLevel int64) error {
    r._cardLevel = _cardLevel
    r.Set("card_level", _cardLevel)
    return nil
}

// CardLevel Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetCardLevel() int64 {
    return r._cardLevel
}
// ConsumeValue Setter
// 核销金额，精确到分，例如1.99元=199
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeValue(_consumeValue int64) error {
    r._consumeValue = _consumeValue
    r.Set("consume_value", _consumeValue)
    return nil
}

// ConsumeValue Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeValue() int64 {
    return r._consumeValue
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Reason Setter
// 核销原因
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetReason() string {
    return r._reason
}
// StoreId Setter
// 门店id
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetStoreId() int64 {
    return r._storeId
}
// OperatorId Setter
// 操作人id
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetOperatorId(_operatorId int64) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetOperatorId() int64 {
    return r._operatorId
}
// ConsumeSerialNum Setter
// 核销流水号，外部ISV全局唯一
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeSerialNum(_consumeSerialNum string) error {
    r._consumeSerialNum = _consumeSerialNum
    r.Set("consume_serial_num", _consumeSerialNum)
    return nil
}

// ConsumeSerialNum Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeSerialNum() string {
    return r._consumeSerialNum
}
// ConsumeCode Setter
// 核销code
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetConsumeCode(_consumeCode string) error {
    r._consumeCode = _consumeCode
    r.Set("consume_code", _consumeCode)
    return nil
}

// ConsumeCode Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetConsumeCode() string {
    return r._consumeCode
}
// Token Setter
// 安全token
func (r *TaobaoVmarketEticketCardConsumecardRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketCardConsumecardRequest) GetToken() string {
    return r._token
}
