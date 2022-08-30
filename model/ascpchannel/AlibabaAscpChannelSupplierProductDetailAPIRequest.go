package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductDetailAPIRequest 供应链渠道中心分销品详情查询(供应商专用) API请求
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
type AlibabaAscpChannelSupplierProductDetailAPIRequest struct {
	model.Params
	// 入参
	_productDetailRequest *ProductDetailQueryRequestForSupplier
}

// NewAlibabaAscpChannelSupplierProductDetailRequest 初始化AlibabaAscpChannelSupplierProductDetailAPIRequest对象
func NewAlibabaAscpChannelSupplierProductDetailRequest() *AlibabaAscpChannelSupplierProductDetailAPIRequest {
	return &AlibabaAscpChannelSupplierProductDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProductDetailRequest is ProductDetailRequest Setter
// 入参
func (r *AlibabaAscpChannelSupplierProductDetailAPIRequest) SetProductDetailRequest(_productDetailRequest *ProductDetailQueryRequestForSupplier) error {
	r._productDetailRequest = _productDetailRequest
	r.Set("product_detail_request", _productDetailRequest)
	return nil
}

// GetProductDetailRequest ProductDetailRequest Getter
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetProductDetailRequest() *ProductDetailQueryRequestForSupplier {
	return r._productDetailRequest
}
