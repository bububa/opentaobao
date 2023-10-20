package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTpFundsRecoverQueryAPIRequest 服务商资金权益逆向扣回的查询接口 API请求
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
type TmallServicecenterTpFundsRecoverQueryAPIRequest struct {
	model.Params
	// query入参
	_query *TpFundsRecoverQuery
}

// NewTmallServicecenterTpFundsRecoverQueryRequest 初始化TmallServicecenterTpFundsRecoverQueryAPIRequest对象
func NewTmallServicecenterTpFundsRecoverQueryRequest() *TmallServicecenterTpFundsRecoverQueryAPIRequest {
	return &TmallServicecenterTpFundsRecoverQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterTpFundsRecoverQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tp.funds.recover.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// query入参
func (r *TmallServicecenterTpFundsRecoverQueryAPIRequest) SetQuery(_query *TpFundsRecoverQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetQuery() *TpFundsRecoverQuery {
	return r._query
}

var poolTmallServicecenterTpFundsRecoverQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterTpFundsRecoverQueryRequest()
	},
}

// GetTmallServicecenterTpFundsRecoverQueryRequest 从 sync.Pool 获取 TmallServicecenterTpFundsRecoverQueryAPIRequest
func GetTmallServicecenterTpFundsRecoverQueryAPIRequest() *TmallServicecenterTpFundsRecoverQueryAPIRequest {
	return poolTmallServicecenterTpFundsRecoverQueryAPIRequest.Get().(*TmallServicecenterTpFundsRecoverQueryAPIRequest)
}

// ReleaseTmallServicecenterTpFundsRecoverQueryAPIRequest 将 TmallServicecenterTpFundsRecoverQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterTpFundsRecoverQueryAPIRequest(v *TmallServicecenterTpFundsRecoverQueryAPIRequest) {
	v.Reset()
	poolTmallServicecenterTpFundsRecoverQueryAPIRequest.Put(v)
}
