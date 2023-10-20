package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosonsitetradequeryAPIRequest 新商场当面付交易查询 API请求
// alibaba.mos.onsite.trade.query
//
// 本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
// 商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。
type AlibabamosonsitetradequeryAPIRequest struct {
	model.Params
	// 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  > 商户支付流水号。
	_tradeNo string
	// 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 >  商户支付流水号。
	_outTradeNo string
	// 商户ID类型，取值为miaojie或out
	_storeIdType string
	// 门店号或喵街商户ID
	_storeId string
}

// NewAlibabamosonsitetradequeryRequest 初始化AlibabamosonsitetradequeryAPIRequest对象
func NewAlibabamosonsitetradequeryRequest() *AlibabamosonsitetradequeryAPIRequest {
	return &AlibabamosonsitetradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosonsitetradequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosonsitetradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosonsitetradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  &gt; 商户支付流水号。
func (r *AlibabamosonsitetradequeryAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabamosonsitetradequeryAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 &gt;  商户支付流水号。
func (r *AlibabamosonsitetradequeryAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabamosonsitetradequeryAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreIdType is StoreIdType Setter
// 商户ID类型，取值为miaojie或out
func (r *AlibabamosonsitetradequeryAPIRequest) SetStoreIdType(_storeIdType string) error {
	r._storeIdType = _storeIdType
	r.Set("store_id_type", _storeIdType)
	return nil
}

// GetStoreIdType StoreIdType Getter
func (r AlibabamosonsitetradequeryAPIRequest) GetStoreIdType() string {
	return r._storeIdType
}

// SetStoreId is StoreId Setter
// 门店号或喵街商户ID
func (r *AlibabamosonsitetradequeryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabamosonsitetradequeryAPIRequest) GetStoreId() string {
	return r._storeId
}
