package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductListAPIRequest 供应链渠道中心淘外分销品批量查询(分销商专用) API请求
// alibaba.ascp.channel.distributor.product.list
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorProductListAPIRequest struct {
	model.Params
	// 列表请求
	_productListRequest *Productlistrequest
}

// NewAlibabaAscpChannelDistributorProductListRequest 初始化AlibabaAscpChannelDistributorProductListAPIRequest对象
func NewAlibabaAscpChannelDistributorProductListRequest() *AlibabaAscpChannelDistributorProductListAPIRequest {
	return &AlibabaAscpChannelDistributorProductListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelDistributorProductListAPIRequest) Reset() {
	r._productListRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductListRequest is ProductListRequest Setter
// 列表请求
func (r *AlibabaAscpChannelDistributorProductListAPIRequest) SetProductListRequest(_productListRequest *Productlistrequest) error {
	r._productListRequest = _productListRequest
	r.Set("product_list_request", _productListRequest)
	return nil
}

// GetProductListRequest ProductListRequest Getter
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetProductListRequest() *Productlistrequest {
	return r._productListRequest
}

var poolAlibabaAscpChannelDistributorProductListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelDistributorProductListRequest()
	},
}

// GetAlibabaAscpChannelDistributorProductListRequest 从 sync.Pool 获取 AlibabaAscpChannelDistributorProductListAPIRequest
func GetAlibabaAscpChannelDistributorProductListAPIRequest() *AlibabaAscpChannelDistributorProductListAPIRequest {
	return poolAlibabaAscpChannelDistributorProductListAPIRequest.Get().(*AlibabaAscpChannelDistributorProductListAPIRequest)
}

// ReleaseAlibabaAscpChannelDistributorProductListAPIRequest 将 AlibabaAscpChannelDistributorProductListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelDistributorProductListAPIRequest(v *AlibabaAscpChannelDistributorProductListAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelDistributorProductListAPIRequest.Put(v)
}
