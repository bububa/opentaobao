package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderGetAPIRequest get order list API请求
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
type AliexpressSolutionOrderGetAPIRequest struct {
	model.Params
	// param
	_param0 *OrderQuery
}

// NewAliexpressSolutionOrderGetRequest 初始化AliexpressSolutionOrderGetAPIRequest对象
func NewAliexpressSolutionOrderGetRequest() *AliexpressSolutionOrderGetAPIRequest {
	return &AliexpressSolutionOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// param
func (r *AliexpressSolutionOrderGetAPIRequest) SetParam0(_param0 *OrderQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AliexpressSolutionOrderGetAPIRequest) GetParam0() *OrderQuery {
	return r._param0
}
