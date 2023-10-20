package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsshippingmethodqueryAPIRequest query shipping method API请求
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
type AliexpresslocallogisticsshippingmethodqueryAPIRequest struct {
	model.Params
	// method query object
	_param1 *QueryShippingMethodRequestTopDto
}

// NewAliexpresslocallogisticsshippingmethodqueryRequest 初始化AliexpresslocallogisticsshippingmethodqueryAPIRequest对象
func NewAliexpresslocallogisticsshippingmethodqueryRequest() *AliexpresslocallogisticsshippingmethodqueryAPIRequest {
	return &AliexpresslocallogisticsshippingmethodqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticsshippingmethodqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.shipping.method.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticsshippingmethodqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticsshippingmethodqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// method query object
func (r *AliexpresslocallogisticsshippingmethodqueryAPIRequest) SetParam1(_param1 *QueryShippingMethodRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticsshippingmethodqueryAPIRequest) GetParam1() *QueryShippingMethodRequestTopDto {
	return r._param1
}
