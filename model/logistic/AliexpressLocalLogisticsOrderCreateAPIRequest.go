package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsordercreateAPIRequest create logistics order API请求
// aliexpress.local.logistics.order.create
//
// create logistics order
type AliexpresslocallogisticsordercreateAPIRequest struct {
	model.Params
	// create logistics order's param
	_param1 *CreateOrderRequestTopDto
}

// NewAliexpresslocallogisticsordercreateRequest 初始化AliexpresslocallogisticsordercreateAPIRequest对象
func NewAliexpresslocallogisticsordercreateRequest() *AliexpresslocallogisticsordercreateAPIRequest {
	return &AliexpresslocallogisticsordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticsordercreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticsordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticsordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// create logistics order&#39;s param
func (r *AliexpresslocallogisticsordercreateAPIRequest) SetParam1(_param1 *CreateOrderRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticsordercreateAPIRequest) GetParam1() *CreateOrderRequestTopDto {
	return r._param1
}
