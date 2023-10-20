package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleautotradeisvorderstateprocessAPIRequest 闲鱼订单状态推进 API请求
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
type AlibabaidleautotradeisvorderstateprocessAPIRequest struct {
	model.Params
	// AutoTradeAPI通用入参
	_autoTradeApiParam *AutoTradeApiparam
}

// NewAlibabaidleautotradeisvorderstateprocessRequest 初始化AlibabaidleautotradeisvorderstateprocessAPIRequest对象
func NewAlibabaidleautotradeisvorderstateprocessRequest() *AlibabaidleautotradeisvorderstateprocessAPIRequest {
	return &AlibabaidleautotradeisvorderstateprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleautotradeisvorderstateprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.autotrade.isv.order.state.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleautotradeisvorderstateprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleautotradeisvorderstateprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAutoTradeApiParam is AutoTradeApiParam Setter
// AutoTradeAPI通用入参
func (r *AlibabaidleautotradeisvorderstateprocessAPIRequest) SetAutoTradeApiParam(_autoTradeApiParam *AutoTradeApiparam) error {
	r._autoTradeApiParam = _autoTradeApiParam
	r.Set("auto_trade_api_param", _autoTradeApiParam)
	return nil
}

// GetAutoTradeApiParam AutoTradeApiParam Getter
func (r AlibabaidleautotradeisvorderstateprocessAPIRequest) GetAutoTradeApiParam() *AutoTradeApiparam {
	return r._autoTradeApiParam
}
