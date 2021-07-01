package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketCardConsumecardAPIRequest
电子凭证储值卡核销 API请求
taobao.vmarket.eticket.card.consumecard

线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额 */
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

// New
