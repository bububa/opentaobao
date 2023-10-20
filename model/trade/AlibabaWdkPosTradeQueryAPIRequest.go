package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradequeryAPIRequest 轻pos品牌营销查询接口 API请求
// alibaba.wdk.pos.trade.query
//
// 轻pos品牌营销场景，外部商家查询营销信息
type AlibabawdkpostradequeryAPIRequest struct {
	model.Params
	// 查询请求
	_queryRequest *FastBuyPosQueryRequest
}

// NewAlibabawdkpostradequeryRequest 初始化AlibabawdkpostradequeryAPIRequest对象
func NewAlibabawdkpostradequeryRequest() *AlibabawdkpostradequeryAPIRequest {
	return &AlibabawdkpostradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpostradequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpostradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpostradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 查询请求
func (r *AlibabawdkpostradequeryAPIRequest) SetQueryRequest(_queryRequest *FastBuyPosQueryRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r AlibabawdkpostradequeryAPIRequest) GetQueryRequest() *FastBuyPosQueryRequest {
	return r._queryRequest
}
