package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsOrderCreateAPIRequest create logistics order API请求
// aliexpress.local.logistics.order.create
//
// create logistics order
type AliexpressLocalLogisticsOrderCreateAPIRequest struct {
	model.Params
	// create logistics order's param
	_param1 *CreateOrderRequestTopDto
}

// NewAliexpressLocalLogisticsOrderCreateRequest 初始化AliexpressLocalLogisticsOrderCreateAPIRequest对象
func NewAliexpressLocalLogisticsOrderCreateRequest() *AliexpressLocalLogisticsOrderCreateAPIRequest {
	return &AliexpressLocalLogisticsOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam1 is Param1 Setter
// create logistics order's param
func (r *AliexpressLocalLogisticsOrderCreateAPIRequest) SetParam1(_param1 *CreateOrderRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetParam1() *CreateOrderRequestTopDto {
	return r._param1
}
