package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTpFundsSendQueryAPIRequest 服务商资金权益发放的查询接口 API请求
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
type TmallServicecenterTpFundsSendQueryAPIRequest struct {
	model.Params
	// 入参对象
	_query *TpFundsSendQuery
}

// NewTmallServicecenterTpFundsSendQueryRequest 初始化TmallServicecenterTpFundsSendQueryAPIRequest对象
func NewTmallServicecenterTpFundsSendQueryRequest() *TmallServicecenterTpFundsSendQueryAPIRequest {
	return &TmallServicecenterTpFundsSendQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterTpFundsSendQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tp.funds.send.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参对象
func (r *TmallServicecenterTpFundsSendQueryAPIRequest) SetQuery(_query *TpFundsSendQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetQuery() *TpFundsSendQuery {
	return r._query
}

var poolTmallServicecenterTpFundsSendQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterTpFundsSendQueryRequest()
	},
}

// GetTmallServicecenterTpFundsSendQueryRequest 从 sync.Pool 获取 TmallServicecenterTpFundsSendQueryAPIRequest
func GetTmallServicecenterTpFundsSendQueryAPIRequest() *TmallServicecenterTpFundsSendQueryAPIRequest {
	return poolTmallServicecenterTpFundsSendQueryAPIRequest.Get().(*TmallServicecenterTpFundsSendQueryAPIRequest)
}

// ReleaseTmallServicecenterTpFundsSendQueryAPIRequest 将 TmallServicecenterTpFundsSendQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterTpFundsSendQueryAPIRequest(v *TmallServicecenterTpFundsSendQueryAPIRequest) {
	v.Reset()
	poolTmallServicecenterTpFundsSendQueryAPIRequest.Put(v)
}
