package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorPriceGetAPIRequest 链渠道中心淘外分销价格查询(分销商专用) API请求
// alibaba.ascp.channel.distributor.price.get
//
// 此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorPriceGetAPIRequest struct {
	model.Params
	// 价格入参
	_priceRequest *Pricerequest
}

// NewAlibabaAscpChannelDistributorPriceGetRequest 初始化AlibabaAscpChannelDistributorPriceGetAPIRequest对象
func NewAlibabaAscpChannelDistributorPriceGetRequest() *AlibabaAscpChannelDistributorPriceGetAPIRequest {
	return &AlibabaAscpChannelDistributorPriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPriceRequest is PriceRequest Setter
// 价格入参
func (r *AlibabaAscpChannelDistributorPriceGetAPIRequest) SetPriceRequest(_priceRequest *Pricerequest) error {
	r._priceRequest = _priceRequest
	r.Set("price_request", _priceRequest)
	return nil
}

// GetPriceRequest PriceRequest Getter
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetPriceRequest() *Pricerequest {
	return r._priceRequest
}
