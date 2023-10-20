package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionordergetAPIRequest get order list API请求
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
type AliexpresssolutionordergetAPIRequest struct {
	model.Params
	// param
	_param0 *OrderQuery
}

// NewAliexpresssolutionordergetRequest 初始化AliexpresssolutionordergetAPIRequest对象
func NewAliexpresssolutionordergetRequest() *AliexpresssolutionordergetAPIRequest {
	return &AliexpresssolutionordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionordergetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// param
func (r *AliexpresssolutionordergetAPIRequest) SetParam0(_param0 *OrderQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AliexpresssolutionordergetAPIRequest) GetParam0() *OrderQuery {
	return r._param0
}
