package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeCancelAPIRequest 零售+平台取消订单 API请求
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
type AlibabaNlifeB2cTradeCancelAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店号
	_storeId string
}

// NewAlibabaNlifeB2cTradeCancelRequest 初始化AlibabaNlifeB2cTradeCancelAPIRequest对象
func NewAlibabaNlifeB2cTradeCancelRequest() *AlibabaNlifeB2cTradeCancelAPIRequest {
	return &AlibabaNlifeB2cTradeCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// Get TradeNo Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// Set is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// Get OutTradeNo Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// Set is StoreId Setter
// 零售+门店号
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetStoreId() string {
	return r._storeId
}
