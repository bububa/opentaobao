package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardQueryAPIRequest 查询交易结果 API请求
// taobao.elife.lifecard.query
//
// 卖家在交易状态不明的情况下, 查询交易结果.
type TaobaoElifeLifecardQueryAPIRequest struct {
	model.Params
	// 入参
	_queryRequest *ConsumeRequest
}

// NewTaobaoElifeLifecardQueryRequest 初始化TaobaoElifeLifecardQueryAPIRequest对象
func NewTaobaoElifeLifecardQueryRequest() *TaobaoElifeLifecardQueryAPIRequest {
	return &TaobaoElifeLifecardQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardQueryAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryRequest Setter
// 入参
func (r *TaobaoElifeLifecardQueryAPIRequest) SetQueryRequest(_queryRequest *ConsumeRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// Get QueryRequest Getter
func (r TaobaoElifeLifecardQueryAPIRequest) GetQueryRequest() *ConsumeRequest {
	return r._queryRequest
}
