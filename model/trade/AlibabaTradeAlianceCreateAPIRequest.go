package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTradeAlianceCreateAPIRequest 推客平台订单回流 API请求
// alibaba.trade.aliance.create
//
// 推客平台订单回流
type AlibabaTradeAlianceCreateAPIRequest struct {
	model.Params
	// 下单请求
	_paramIsvCreateOrderParam *IsvCreateOrderParam
}

// NewAlibabaTradeAlianceCreateRequest 初始化AlibabaTradeAlianceCreateAPIRequest对象
func NewAlibabaTradeAlianceCreateRequest() *AlibabaTradeAlianceCreateAPIRequest {
	return &AlibabaTradeAlianceCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTradeAlianceCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.trade.aliance.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTradeAlianceCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTradeAlianceCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIsvCreateOrderParam is ParamIsvCreateOrderParam Setter
// 下单请求
func (r *AlibabaTradeAlianceCreateAPIRequest) SetParamIsvCreateOrderParam(_paramIsvCreateOrderParam *IsvCreateOrderParam) error {
	r._paramIsvCreateOrderParam = _paramIsvCreateOrderParam
	r.Set("param_isv_create_order_param", _paramIsvCreateOrderParam)
	return nil
}

// GetParamIsvCreateOrderParam ParamIsvCreateOrderParam Getter
func (r AlibabaTradeAlianceCreateAPIRequest) GetParamIsvCreateOrderParam() *IsvCreateOrderParam {
	return r._paramIsvCreateOrderParam
}
