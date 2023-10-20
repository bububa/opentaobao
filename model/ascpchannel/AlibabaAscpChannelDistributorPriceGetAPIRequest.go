package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorpricegetAPIRequest 链渠道中心淘外分销价格查询(分销商专用) API请求
// alibaba.ascp.channel.distributor.price.get
//
// 此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
type AlibabaascpchanneldistributorpricegetAPIRequest struct {
	model.Params
	// 价格入参
	_priceRequest *Pricerequest
}

// NewAlibabaascpchanneldistributorpricegetRequest 初始化AlibabaascpchanneldistributorpricegetAPIRequest对象
func NewAlibabaascpchanneldistributorpricegetRequest() *AlibabaascpchanneldistributorpricegetAPIRequest {
	return &AlibabaascpchanneldistributorpricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorpricegetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorpricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorpricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriceRequest is PriceRequest Setter
// 价格入参
func (r *AlibabaascpchanneldistributorpricegetAPIRequest) SetPriceRequest(_priceRequest *Pricerequest) error {
	r._priceRequest = _priceRequest
	r.Set("price_request", _priceRequest)
	return nil
}

// GetPriceRequest PriceRequest Getter
func (r AlibabaascpchanneldistributorpricegetAPIRequest) GetPriceRequest() *Pricerequest {
	return r._priceRequest
}
