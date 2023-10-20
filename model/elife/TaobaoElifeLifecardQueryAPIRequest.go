package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoelifelifecardqueryAPIRequest 查询交易结果 API请求
// taobao.elife.lifecard.query
//
// 卖家在交易状态不明的情况下, 查询交易结果.
type TaobaoelifelifecardqueryAPIRequest struct {
	model.Params
	// 入参
	_queryRequest *ConsumeRequest
}

// NewTaobaoelifelifecardqueryRequest 初始化TaobaoelifelifecardqueryAPIRequest对象
func NewTaobaoelifelifecardqueryRequest() *TaobaoelifelifecardqueryAPIRequest {
	return &TaobaoelifelifecardqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoelifelifecardqueryAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoelifelifecardqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoelifelifecardqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 入参
func (r *TaobaoelifelifecardqueryAPIRequest) SetQueryRequest(_queryRequest *ConsumeRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r TaobaoelifelifecardqueryAPIRequest) GetQueryRequest() *ConsumeRequest {
	return r._queryRequest
}
