package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
