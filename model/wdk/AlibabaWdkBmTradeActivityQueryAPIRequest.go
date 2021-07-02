package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmTradeActivityQueryAPIRequest 品牌营销的订单活动信息查询 API请求
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
type AlibabaWdkBmTradeActivityQueryAPIRequest struct {
	model.Params
	// 入参
	_queryParam *IsvOrderQueryParam
}

// NewAlibabaWdkBmTradeActivityQueryRequest 初始化AlibabaWdkBmTradeActivityQueryAPIRequest对象
func NewAlibabaWdkBmTradeActivityQueryRequest() *AlibabaWdkBmTradeActivityQueryAPIRequest {
	return &AlibabaWdkBmTradeActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.trade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryParam is QueryParam Setter
// 入参
func (r *AlibabaWdkBmTradeActivityQueryAPIRequest) SetQueryParam(_queryParam *IsvOrderQueryParam) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetQueryParam() *IsvOrderQueryParam {
	return r._queryParam
}
