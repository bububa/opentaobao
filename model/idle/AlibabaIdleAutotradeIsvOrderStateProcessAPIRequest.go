package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest 闲鱼订单状态推进 API请求
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
type AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest struct {
	model.Params
	// AutoTradeAPI通用入参
	_autoTradeApiParam *AutoTradeApiparam
}

// NewAlibabaIdleAutotradeIsvOrderStateProcessRequest 初始化AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest对象
func NewAlibabaIdleAutotradeIsvOrderStateProcessRequest() *AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest {
	return &AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.autotrade.isv.order.state.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAutoTradeApiParam is AutoTradeApiParam Setter
// AutoTradeAPI通用入参
func (r *AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest) SetAutoTradeApiParam(_autoTradeApiParam *AutoTradeApiparam) error {
	r._autoTradeApiParam = _autoTradeApiParam
	r.Set("auto_trade_api_param", _autoTradeApiParam)
	return nil
}

// GetAutoTradeApiParam AutoTradeApiParam Getter
func (r AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest) GetAutoTradeApiParam() *AutoTradeApiparam {
	return r._autoTradeApiParam
}
