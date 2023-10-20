package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctradegetAPIRequest 零售+平台查询订单 API请求
// alibaba.nlife.b2c.trade.get
//
// 查询零售+平台创建出来的订单详情
type Alibabanlifeb2ctradegetAPIRequest struct {
	model.Params
	// 零售+平台订单号,和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
	_storeId string
}

// NewAlibabanlifeb2ctradegetRequest 初始化Alibabanlifeb2ctradegetAPIRequest对象
func NewAlibabanlifeb2ctradegetRequest() *Alibabanlifeb2ctradegetAPIRequest {
	return &Alibabanlifeb2ctradegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2ctradegetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2ctradegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2ctradegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号,和out_trade_no不能同时为空
func (r *Alibabanlifeb2ctradegetAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r Alibabanlifeb2ctradegetAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *Alibabanlifeb2ctradegetAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r Alibabanlifeb2ctradegetAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreId is StoreId Setter
// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
func (r *Alibabanlifeb2ctradegetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2ctradegetAPIRequest) GetStoreId() string {
	return r._storeId
}
