package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctradecancelAPIRequest 零售+平台取消订单 API请求
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
type Alibabanlifeb2ctradecancelAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店号
	_storeId string
}

// NewAlibabanlifeb2ctradecancelRequest 初始化Alibabanlifeb2ctradecancelAPIRequest对象
func NewAlibabanlifeb2ctradecancelRequest() *Alibabanlifeb2ctradecancelAPIRequest {
	return &Alibabanlifeb2ctradecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2ctradecancelAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2ctradecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2ctradecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *Alibabanlifeb2ctradecancelAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r Alibabanlifeb2ctradecancelAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *Alibabanlifeb2ctradecancelAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r Alibabanlifeb2ctradecancelAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreId is StoreId Setter
// 零售+门店号
func (r *Alibabanlifeb2ctradecancelAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2ctradecancelAPIRequest) GetStoreId() string {
	return r._storeId
}
