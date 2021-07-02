package servicecenter

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tp.funds.recover.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsRecoverQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
