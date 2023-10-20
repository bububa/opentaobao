package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsorderinfoqueryAPIRequest query order details API请求
// aliexpress.local.logistics.order.info.query
//
// query order details
type AliexpresslocallogisticsorderinfoqueryAPIRequest struct {
	model.Params
	// param1
	_param1 *FindOrderRequestTopDto
}

// NewAliexpresslocallogisticsorderinfoqueryRequest 初始化AliexpresslocallogisticsorderinfoqueryAPIRequest对象
func NewAliexpresslocallogisticsorderinfoqueryRequest() *AliexpresslocallogisticsorderinfoqueryAPIRequest {
	return &AliexpresslocallogisticsorderinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticsorderinfoqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.order.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticsorderinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticsorderinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param1
func (r *AliexpresslocallogisticsorderinfoqueryAPIRequest) SetParam1(_param1 *FindOrderRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticsorderinfoqueryAPIRequest) GetParam1() *FindOrderRequestTopDto {
	return r._param1
}
