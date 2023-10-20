package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductdetailAPIRequest 供应链渠道中心分销品详情查询(供应商专用) API请求
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
type AlibabaascpchannelsupplierproductdetailAPIRequest struct {
	model.Params
	// 入参
	_productDetailRequest *ProductDetailQueryRequestForSupplier
}

// NewAlibabaascpchannelsupplierproductdetailRequest 初始化AlibabaascpchannelsupplierproductdetailAPIRequest对象
func NewAlibabaascpchannelsupplierproductdetailRequest() *AlibabaascpchannelsupplierproductdetailAPIRequest {
	return &AlibabaascpchannelsupplierproductdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsupplierproductdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsupplierproductdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsupplierproductdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductDetailRequest is ProductDetailRequest Setter
// 入参
func (r *AlibabaascpchannelsupplierproductdetailAPIRequest) SetProductDetailRequest(_productDetailRequest *ProductDetailQueryRequestForSupplier) error {
	r._productDetailRequest = _productDetailRequest
	r.Set("product_detail_request", _productDetailRequest)
	return nil
}

// GetProductDetailRequest ProductDetailRequest Getter
func (r AlibabaascpchannelsupplierproductdetailAPIRequest) GetProductDetailRequest() *ProductDetailQueryRequestForSupplier {
	return r._productDetailRequest
}
