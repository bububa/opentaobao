package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertpfundsrecoverqueryAPIRequest 服务商资金权益逆向扣回的查询接口 API请求
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
type TmallservicecentertpfundsrecoverqueryAPIRequest struct {
	model.Params
	// query入参
	_query *TpFundsRecoverQuery
}

// NewTmallservicecentertpfundsrecoverqueryRequest 初始化TmallservicecentertpfundsrecoverqueryAPIRequest对象
func NewTmallservicecentertpfundsrecoverqueryRequest() *TmallservicecentertpfundsrecoverqueryAPIRequest {
	return &TmallservicecentertpfundsrecoverqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertpfundsrecoverqueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tp.funds.recover.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertpfundsrecoverqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertpfundsrecoverqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// query入参
func (r *TmallservicecentertpfundsrecoverqueryAPIRequest) SetQuery(_query *TpFundsRecoverQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallservicecentertpfundsrecoverqueryAPIRequest) GetQuery() *TpFundsRecoverQuery {
	return r._query
}
