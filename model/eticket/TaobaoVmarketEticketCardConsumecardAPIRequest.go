package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketCardConsumecardAPIRequest 电子凭证储值卡核销 API请求
// taobao.vmarket.eticket.card.consumecard
//
// 线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
type TaobaoVmarketEticketCardConsumecardAPIRequest struct {
	model.Params
	// 卡号
	_cardId string
	// 卡内等级
	_cardLevel int64
	// 核销金额，精确到分，例如1.99元=199
	_consumeValue int64
	// 买家昵称
	_buyerNick string
	// 核销原因
	_reason string
	// 门店id
	_storeId int64
	// 操作人id
	_operatorId int64
	// 核销流水号，外部ISV全局唯一
	_consumeSerialNum string
	// 核销code
	_consumeCode string
	// 安全token
	_token string
}

// NewTaobaoVmarketEticketCardConsumecardRequest 初始化TaobaoVmarketEticketCardConsumecardAPIRequest对象
func NewTaobaoVmarketEticketCardConsumecardRequest() *TaobaoVmarketEticketCardConsumecardAPIRequest {
	return &TaobaoVmarketEticketCardConsumecardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.card.consumecard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCardId is CardId Setter
// 卡号
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetCardId(_cardId string) error {
	r._cardId = _cardId
	r.Set("card_id", _cardId)
	return nil
}

// GetCardId CardId Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetCardId() string {
	return r._cardId
}

// SetCardLevel is CardLevel Setter
// 卡内等级
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetCardLevel(_cardLevel int64) error {
	r._cardLevel = _cardLevel
	r.Set("card_level", _cardLevel)
	return nil
}

// GetCardLevel CardLevel Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetCardLevel() int64 {
	return r._cardLevel
}

// SetConsumeValue is ConsumeValue Setter
// 核销金额，精确到分，例如1.99元=199
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetConsumeValue(_consumeValue int64) error {
	r._consumeValue = _consumeValue
	r.Set("consume_value", _consumeValue)
	return nil
}

// GetConsumeValue ConsumeValue Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetConsumeValue() int64 {
	return r._consumeValue
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetReason is Reason Setter
// 核销原因
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetReason() string {
	return r._reason
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetOperatorId is OperatorId Setter
// 操作人id
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetOperatorId(_operatorId int64) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// GetOperatorId OperatorId Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetOperatorId() int64 {
	return r._operatorId
}

// SetConsumeSerialNum is ConsumeSerialNum Setter
// 核销流水号，外部ISV全局唯一
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetConsumeSerialNum(_consumeSerialNum string) error {
	r._consumeSerialNum = _consumeSerialNum
	r.Set("consume_serial_num", _consumeSerialNum)
	return nil
}

// GetConsumeSerialNum ConsumeSerialNum Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetConsumeSerialNum() string {
	return r._consumeSerialNum
}

// SetConsumeCode is ConsumeCode Setter
// 核销code
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetConsumeCode(_consumeCode string) error {
	r._consumeCode = _consumeCode
	r.Set("consume_code", _consumeCode)
	return nil
}

// GetConsumeCode ConsumeCode Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetConsumeCode() string {
	return r._consumeCode
}

// SetToken is Token Setter
// 安全token
func (r *TaobaoVmarketEticketCardConsumecardAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoVmarketEticketCardConsumecardAPIRequest) GetToken() string {
	return r._token
}
