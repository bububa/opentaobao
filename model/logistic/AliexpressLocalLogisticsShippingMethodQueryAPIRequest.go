package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsShippingMethodQueryAPIRequest query shipping method API请求
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
type AliexpressLocalLogisticsShippingMethodQueryAPIRequest struct {
	model.Params
	// method query object
	_param1 *QueryShippingMethodRequestTopDto
}

// NewAliexpressLocalLogisticsShippingMethodQueryRequest 初始化AliexpressLocalLogisticsShippingMethodQueryAPIRequest对象
func NewAliexpressLocalLogisticsShippingMethodQueryRequest() *AliexpressLocalLogisticsShippingMethodQueryAPIRequest {
	return &AliexpressLocalLogisticsShippingMethodQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsShippingMethodQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.shipping.method.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsShippingMethodQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam1 is Param1 Setter
// method query object
func (r *AliexpressLocalLogisticsShippingMethodQueryAPIRequest) SetParam1(_param1 *QueryShippingMethodRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsShippingMethodQueryAPIRequest) GetParam1() *QueryShippingMethodRequestTopDto {
	return r._param1
}
