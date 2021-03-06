package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderInfoGetAPIRequest get order detail info API请求
// aliexpress.solution.order.info.get
//
// get order detail info
type AliexpressSolutionOrderInfoGetAPIRequest struct {
	model.Params
	// param
	_param1 *OrderDetailQuery
}

// NewAliexpressSolutionOrderInfoGetRequest 初始化AliexpressSolutionOrderInfoGetAPIRequest对象
func NewAliexpressSolutionOrderInfoGetRequest() *AliexpressSolutionOrderInfoGetAPIRequest {
	return &AliexpressSolutionOrderInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpressSolutionOrderInfoGetAPIRequest) SetParam1(_param1 *OrderDetailQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetParam1() *OrderDetailQuery {
	return r._param1
}
