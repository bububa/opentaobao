package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelDistributorPriceGetAPIRequest) Reset() {
	r._priceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorPriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpChannelDistributorPriceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelDistributorPriceGetRequest()
	},
}

// GetAlibabaAscpChannelDistributorPriceGetRequest 从 sync.Pool 获取 AlibabaAscpChannelDistributorPriceGetAPIRequest
func GetAlibabaAscpChannelDistributorPriceGetAPIRequest() *AlibabaAscpChannelDistributorPriceGetAPIRequest {
	return poolAlibabaAscpChannelDistributorPriceGetAPIRequest.Get().(*AlibabaAscpChannelDistributorPriceGetAPIRequest)
}

// ReleaseAlibabaAscpChannelDistributorPriceGetAPIRequest 将 AlibabaAscpChannelDistributorPriceGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelDistributorPriceGetAPIRequest(v *AlibabaAscpChannelDistributorPriceGetAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelDistributorPriceGetAPIRequest.Put(v)
}
