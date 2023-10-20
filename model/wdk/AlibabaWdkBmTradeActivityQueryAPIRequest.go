package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmtradeactivityqueryAPIRequest 品牌营销的订单活动信息查询 API请求
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
type AlibabawdkbmtradeactivityqueryAPIRequest struct {
	model.Params
	// 入参
	_queryParam *IsvOrderQueryParam
}

// NewAlibabawdkbmtradeactivityqueryRequest 初始化AlibabawdkbmtradeactivityqueryAPIRequest对象
func NewAlibabawdkbmtradeactivityqueryRequest() *AlibabawdkbmtradeactivityqueryAPIRequest {
	return &AlibabawdkbmtradeactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkbmtradeactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.trade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkbmtradeactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkbmtradeactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 入参
func (r *AlibabawdkbmtradeactivityqueryAPIRequest) SetQueryParam(_queryParam *IsvOrderQueryParam) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlibabawdkbmtradeactivityqueryAPIRequest) GetQueryParam() *IsvOrderQueryParam {
	return r._queryParam
}
