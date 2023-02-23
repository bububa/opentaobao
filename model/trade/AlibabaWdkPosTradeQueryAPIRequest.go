package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeQueryAPIRequest 轻pos品牌营销查询接口 API请求
// alibaba.wdk.pos.trade.query
//
// 轻pos品牌营销场景，外部商家查询营销信息
type AlibabaWdkPosTradeQueryAPIRequest struct {
	model.Params
	// 查询请求
	_queryRequest *FastBuyPosQueryRequest
}

// NewAlibabaWdkPosTradeQueryRequest 初始化AlibabaWdkPosTradeQueryAPIRequest对象
func NewAlibabaWdkPosTradeQueryRequest() *AlibabaWdkPosTradeQueryAPIRequest {
	return &AlibabaWdkPosTradeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosTradeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 查询请求
func (r *AlibabaWdkPosTradeQueryAPIRequest) SetQueryRequest(_queryRequest *FastBuyPosQueryRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r AlibabaWdkPosTradeQueryAPIRequest) GetQueryRequest() *FastBuyPosQueryRequest {
	return r._queryRequest
}
