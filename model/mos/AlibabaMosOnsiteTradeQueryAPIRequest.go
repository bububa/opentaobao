package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradeQueryAPIRequest
新商场当面付交易查询 API请求
alibaba.mos.onsite.trade.query

本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。 */
type AlibabaMosOnsiteTradeQueryAPIRequest struct {
	model.Params
	// 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  > 商户支付流水号。
	_tradeNo string
	// 商户ID类型，取值为miaojie或out
	_storeIdType string
	// 门店号或喵街商户ID
	_storeId string
	// 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 >  商户支付流水号。
	_outTradeNo string
}

// New
