package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingProductGetAPIRequest 阿里巴巴dropshipping 产品信息获取 API请求
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
type AlibabaDropshippingProductGetAPIRequest struct {
	model.Params
	// {}
	_paramDistributionSaleProductRequest *DistributionSaleProductRequest
}

// NewAlibabaDropshippingProductGetRequest 初始化AlibabaDropshippingProductGetAPIRequest对象
func NewAlibabaDropshippingProductGetRequest() *AlibabaDropshippingProductGetAPIRequest {
	return &AlibabaDropshippingProductGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDropshippingProductGetAPIRequest) Reset() {
	r._paramDistributionSaleProductRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingProductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDropshippingProductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDistributionSaleProductRequest is ParamDistributionSaleProductRequest Setter
// {}
func (r *AlibabaDropshippingProductGetAPIRequest) SetParamDistributionSaleProductRequest(_paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
	r._paramDistributionSaleProductRequest = _paramDistributionSaleProductRequest
	r.Set("param_distribution_sale_product_request", _paramDistributionSaleProductRequest)
	return nil
}

// GetParamDistributionSaleProductRequest ParamDistributionSaleProductRequest Getter
func (r AlibabaDropshippingProductGetAPIRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
	return r._paramDistributionSaleProductRequest
}

var poolAlibabaDropshippingProductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDropshippingProductGetRequest()
	},
}

// GetAlibabaDropshippingProductGetRequest 从 sync.Pool 获取 AlibabaDropshippingProductGetAPIRequest
func GetAlibabaDropshippingProductGetAPIRequest() *AlibabaDropshippingProductGetAPIRequest {
	return poolAlibabaDropshippingProductGetAPIRequest.Get().(*AlibabaDropshippingProductGetAPIRequest)
}

// ReleaseAlibabaDropshippingProductGetAPIRequest 将 AlibabaDropshippingProductGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaDropshippingProductGetAPIRequest(v *AlibabaDropshippingProductGetAPIRequest) {
	v.Reset()
	poolAlibabaDropshippingProductGetAPIRequest.Put(v)
}
