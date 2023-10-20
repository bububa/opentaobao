package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatradealiancecreateAPIRequest 推客平台订单回流 API请求
// alibaba.trade.aliance.create
//
// 推客平台订单回流
type AlibabatradealiancecreateAPIRequest struct {
	model.Params
	// 下单请求
	_paramIsvCreateOrderParam *IsvCreateOrderParam
}

// NewAlibabatradealiancecreateRequest 初始化AlibabatradealiancecreateAPIRequest对象
func NewAlibabatradealiancecreateRequest() *AlibabatradealiancecreateAPIRequest {
	return &AlibabatradealiancecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatradealiancecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.trade.aliance.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatradealiancecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatradealiancecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIsvCreateOrderParam is ParamIsvCreateOrderParam Setter
// 下单请求
func (r *AlibabatradealiancecreateAPIRequest) SetParamIsvCreateOrderParam(_paramIsvCreateOrderParam *IsvCreateOrderParam) error {
	r._paramIsvCreateOrderParam = _paramIsvCreateOrderParam
	r.Set("param_isv_create_order_param", _paramIsvCreateOrderParam)
	return nil
}

// GetParamIsvCreateOrderParam ParamIsvCreateOrderParam Getter
func (r AlibabatradealiancecreateAPIRequest) GetParamIsvCreateOrderParam() *IsvCreateOrderParam {
	return r._paramIsvCreateOrderParam
}
