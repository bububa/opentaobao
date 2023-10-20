package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) Reset() {
	r._advertiseInfoQuery = nil
	r.Params.ToZero()
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

var poolAlibabaAlicomTradeAdvertiseinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomTradeAdvertiseinfoGetRequest()
	},
}

// GetAlibabaAlicomTradeAdvertiseinfoGetRequest 从 sync.Pool 获取 AlibabaAlicomTradeAdvertiseinfoGetAPIRequest
func GetAlibabaAlicomTradeAdvertiseinfoGetAPIRequest() *AlibabaAlicomTradeAdvertiseinfoGetAPIRequest {
	return poolAlibabaAlicomTradeAdvertiseinfoGetAPIRequest.Get().(*AlibabaAlicomTradeAdvertiseinfoGetAPIRequest)
}

// ReleaseAlibabaAlicomTradeAdvertiseinfoGetAPIRequest 将 AlibabaAlicomTradeAdvertiseinfoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomTradeAdvertiseinfoGetAPIRequest(v *AlibabaAlicomTradeAdvertiseinfoGetAPIRequest) {
	v.Reset()
	poolAlibabaAlicomTradeAdvertiseinfoGetAPIRequest.Put(v)
}
