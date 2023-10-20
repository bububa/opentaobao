package elife

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoElifeLifecardQueryAPIRequest) Reset() {
	r._queryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardQueryAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoElifeLifecardQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 入参
func (r *TaobaoElifeLifecardQueryAPIRequest) SetQueryRequest(_queryRequest *ConsumeRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r TaobaoElifeLifecardQueryAPIRequest) GetQueryRequest() *ConsumeRequest {
	return r._queryRequest
}

var poolTaobaoElifeLifecardQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoElifeLifecardQueryRequest()
	},
}

// GetTaobaoElifeLifecardQueryRequest 从 sync.Pool 获取 TaobaoElifeLifecardQueryAPIRequest
func GetTaobaoElifeLifecardQueryAPIRequest() *TaobaoElifeLifecardQueryAPIRequest {
	return poolTaobaoElifeLifecardQueryAPIRequest.Get().(*TaobaoElifeLifecardQueryAPIRequest)
}

// ReleaseTaobaoElifeLifecardQueryAPIRequest 将 TaobaoElifeLifecardQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoElifeLifecardQueryAPIRequest(v *TaobaoElifeLifecardQueryAPIRequest) {
	v.Reset()
	poolTaobaoElifeLifecardQueryAPIRequest.Put(v)
}
