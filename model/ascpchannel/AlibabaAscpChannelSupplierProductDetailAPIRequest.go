package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSupplierProductDetailAPIRequest) Reset() {
	r._productDetailRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSupplierProductDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpChannelSupplierProductDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSupplierProductDetailRequest()
	},
}

// GetAlibabaAscpChannelSupplierProductDetailRequest 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductDetailAPIRequest
func GetAlibabaAscpChannelSupplierProductDetailAPIRequest() *AlibabaAscpChannelSupplierProductDetailAPIRequest {
	return poolAlibabaAscpChannelSupplierProductDetailAPIRequest.Get().(*AlibabaAscpChannelSupplierProductDetailAPIRequest)
}

// ReleaseAlibabaAscpChannelSupplierProductDetailAPIRequest 将 AlibabaAscpChannelSupplierProductDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductDetailAPIRequest(v *AlibabaAscpChannelSupplierProductDetailAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductDetailAPIRequest.Put(v)
}
