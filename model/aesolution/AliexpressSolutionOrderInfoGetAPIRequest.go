package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionorderinfogetAPIRequest get order detail info API请求
// aliexpress.solution.order.info.get
//
// get order detail info
type AliexpresssolutionorderinfogetAPIRequest struct {
	model.Params
	// param
	_param1 *OrderDetailQuery
}

// NewAliexpresssolutionorderinfogetRequest 初始化AliexpresssolutionorderinfogetAPIRequest对象
func NewAliexpresssolutionorderinfogetRequest() *AliexpresssolutionorderinfogetAPIRequest {
	return &AliexpresssolutionorderinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionorderinfogetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionorderinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionorderinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpresssolutionorderinfogetAPIRequest) SetParam1(_param1 *OrderDetailQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresssolutionorderinfogetAPIRequest) GetParam1() *OrderDetailQuery {
	return r._param1
}
