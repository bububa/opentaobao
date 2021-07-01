package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDropshippingProductGetAPIRequest
阿里巴巴dropshipping 产品信息获取 API请求
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取 */
type AlibabaDropshippingProductGetAPIRequest struct {
	model.Params
	// {}
	_paramDistributionSaleProductRequest *DistributionSaleProductRequest
}

// NewAlibabaDropshippingProductGetRequest 初始化AlibabaDropshippingProductGetAPIRequest对象
func NewAlibabaDropshippingProductGetRequest() *AlibabaDropshippingProductGetAPIRequest {
	return &AlibabaDropshippingProductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingProductGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamDistributionSaleProductRequest Setter
// {}
func (r *AlibabaDropshippingProductGetAPIRequest) SetParamDistributionSaleProductRequest(_paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
	r._paramDistributionSaleProductRequest = _paramDistributionSaleProductRequest
	r.Set("param_distribution_sale_product_request", _paramDistributionSaleProductRequest)
	return nil
}

// Get ParamDistributionSaleProductRequest Getter
func (r AlibabaDropshippingProductGetAPIRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
	return r._paramDistributionSaleProductRequest
}
