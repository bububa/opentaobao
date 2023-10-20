package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductDetailAPIRequest 获取供应链渠道中心品的详情接口（淘外分销商专用） API请求
// alibaba.ascp.channel.distributor.product.detail
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorProductDetailAPIRequest struct {
	model.Params
	// 产品详情查询入参
	_productDetailRequest *ProductDetailQueryRequestForDistributor
}

// NewAlibabaAscpChannelDistributorProductDetailRequest 初始化AlibabaAscpChannelDistributorProductDetailAPIRequest对象
func NewAlibabaAscpChannelDistributorProductDetailRequest() *AlibabaAscpChannelDistributorProductDetailAPIRequest {
	return &AlibabaAscpChannelDistributorProductDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelDistributorProductDetailAPIRequest) Reset() {
	r._productDetailRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductDetailRequest is ProductDetailRequest Setter
// 产品详情查询入参
func (r *AlibabaAscpChannelDistributorProductDetailAPIRequest) SetProductDetailRequest(_productDetailRequest *ProductDetailQueryRequestForDistributor) error {
	r._productDetailRequest = _productDetailRequest
	r.Set("product_detail_request", _productDetailRequest)
	return nil
}

// GetProductDetailRequest ProductDetailRequest Getter
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetProductDetailRequest() *ProductDetailQueryRequestForDistributor {
	return r._productDetailRequest
}

var poolAlibabaAscpChannelDistributorProductDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelDistributorProductDetailRequest()
	},
}

// GetAlibabaAscpChannelDistributorProductDetailRequest 从 sync.Pool 获取 AlibabaAscpChannelDistributorProductDetailAPIRequest
func GetAlibabaAscpChannelDistributorProductDetailAPIRequest() *AlibabaAscpChannelDistributorProductDetailAPIRequest {
	return poolAlibabaAscpChannelDistributorProductDetailAPIRequest.Get().(*AlibabaAscpChannelDistributorProductDetailAPIRequest)
}

// ReleaseAlibabaAscpChannelDistributorProductDetailAPIRequest 将 AlibabaAscpChannelDistributorProductDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelDistributorProductDetailAPIRequest(v *AlibabaAscpChannelDistributorProductDetailAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelDistributorProductDetailAPIRequest.Put(v)
}
