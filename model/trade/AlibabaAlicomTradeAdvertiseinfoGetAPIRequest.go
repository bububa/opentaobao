package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomtradeadvertiseinfogetAPIRequest 获取订单上的在信息流投放信息 API请求
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
type AlibabaalicomtradeadvertiseinfogetAPIRequest struct {
	model.Params
	// 入参
	_advertiseInfoQuery *AdvertiseInfoQuery
}

// NewAlibabaalicomtradeadvertiseinfogetRequest 初始化AlibabaalicomtradeadvertiseinfogetAPIRequest对象
func NewAlibabaalicomtradeadvertiseinfogetRequest() *AlibabaalicomtradeadvertiseinfogetAPIRequest {
	return &AlibabaalicomtradeadvertiseinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomtradeadvertiseinfogetAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.trade.advertiseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomtradeadvertiseinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomtradeadvertiseinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdvertiseInfoQuery is AdvertiseInfoQuery Setter
// 入参
func (r *AlibabaalicomtradeadvertiseinfogetAPIRequest) SetAdvertiseInfoQuery(_advertiseInfoQuery *AdvertiseInfoQuery) error {
	r._advertiseInfoQuery = _advertiseInfoQuery
	r.Set("advertise_info_query", _advertiseInfoQuery)
	return nil
}

// GetAdvertiseInfoQuery AdvertiseInfoQuery Getter
func (r AlibabaalicomtradeadvertiseinfogetAPIRequest) GetAdvertiseInfoQuery() *AdvertiseInfoQuery {
	return r._advertiseInfoQuery
}
