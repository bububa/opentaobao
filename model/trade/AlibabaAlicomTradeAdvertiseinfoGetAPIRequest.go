package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomTradeAdvertiseinfoGetAPIRequest 获取订单上的在信息流投放信息 API请求
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
type AlibabaAlicomTradeAdvertiseinfoGetAPIRequest struct {
	model.Params
	// 入参
	_advertiseInfoQuery *AdvertiseInfoQuery
}

// NewAlibabaAlicomTradeAdvertiseinfoGetRequest 初始化AlibabaAlicomTradeAdvertiseinfoGetAPIRequest对象
func NewAlibabaAlicomTradeAdvertiseinfoGetRequest() *AlibabaAlicomTradeAdvertiseinfoGetAPIRequest {
	return &AlibabaAlicomTradeAdvertiseinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.trade.advertiseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdvertiseInfoQuery is AdvertiseInfoQuery Setter
// 入参
func (r *AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) SetAdvertiseInfoQuery(_advertiseInfoQuery *AdvertiseInfoQuery) error {
	r._advertiseInfoQuery = _advertiseInfoQuery
	r.Set("advertise_info_query", _advertiseInfoQuery)
	return nil
}

// GetAdvertiseInfoQuery AdvertiseInfoQuery Getter
func (r AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) GetAdvertiseInfoQuery() *AdvertiseInfoQuery {
	return r._advertiseInfoQuery
}
