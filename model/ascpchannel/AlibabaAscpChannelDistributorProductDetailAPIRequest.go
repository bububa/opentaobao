package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorproductdetailAPIRequest 获取供应链渠道中心品的详情接口（淘外分销商专用） API请求
// alibaba.ascp.channel.distributor.product.detail
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
type AlibabaascpchanneldistributorproductdetailAPIRequest struct {
	model.Params
	// 产品详情查询入参
	_productDetailRequest *ProductDetailQueryRequestForDistributor
}

// NewAlibabaascpchanneldistributorproductdetailRequest 初始化AlibabaascpchanneldistributorproductdetailAPIRequest对象
func NewAlibabaascpchanneldistributorproductdetailRequest() *AlibabaascpchanneldistributorproductdetailAPIRequest {
	return &AlibabaascpchanneldistributorproductdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorproductdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorproductdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorproductdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductDetailRequest is ProductDetailRequest Setter
// 产品详情查询入参
func (r *AlibabaascpchanneldistributorproductdetailAPIRequest) SetProductDetailRequest(_productDetailRequest *ProductDetailQueryRequestForDistributor) error {
	r._productDetailRequest = _productDetailRequest
	r.Set("product_detail_request", _productDetailRequest)
	return nil
}

// GetProductDetailRequest ProductDetailRequest Getter
func (r AlibabaascpchanneldistributorproductdetailAPIRequest) GetProductDetailRequest() *ProductDetailQueryRequestForDistributor {
	return r._productDetailRequest
}
