package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadropshippingproductgetAPIRequest 阿里巴巴dropshipping 产品信息获取 API请求
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
type AlibabadropshippingproductgetAPIRequest struct {
	model.Params
	// {}
	_paramDistributionSaleProductRequest *DistributionSaleProductRequest
}

// NewAlibabadropshippingproductgetRequest 初始化AlibabadropshippingproductgetAPIRequest对象
func NewAlibabadropshippingproductgetRequest() *AlibabadropshippingproductgetAPIRequest {
	return &AlibabadropshippingproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadropshippingproductgetAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadropshippingproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadropshippingproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDistributionSaleProductRequest is ParamDistributionSaleProductRequest Setter
// {}
func (r *AlibabadropshippingproductgetAPIRequest) SetParamDistributionSaleProductRequest(_paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
	r._paramDistributionSaleProductRequest = _paramDistributionSaleProductRequest
	r.Set("param_distribution_sale_product_request", _paramDistributionSaleProductRequest)
	return nil
}

// GetParamDistributionSaleProductRequest ParamDistributionSaleProductRequest Getter
func (r AlibabadropshippingproductgetAPIRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
	return r._paramDistributionSaleProductRequest
}
