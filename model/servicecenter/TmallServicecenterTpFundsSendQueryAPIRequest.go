package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertpfundssendqueryAPIRequest 服务商资金权益发放的查询接口 API请求
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
type TmallservicecentertpfundssendqueryAPIRequest struct {
	model.Params
	// 入参对象
	_query *TpFundsSendQuery
}

// NewTmallservicecentertpfundssendqueryRequest 初始化TmallservicecentertpfundssendqueryAPIRequest对象
func NewTmallservicecentertpfundssendqueryRequest() *TmallservicecentertpfundssendqueryAPIRequest {
	return &TmallservicecentertpfundssendqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertpfundssendqueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tp.funds.send.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertpfundssendqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertpfundssendqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参对象
func (r *TmallservicecentertpfundssendqueryAPIRequest) SetQuery(_query *TpFundsSendQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallservicecentertpfundssendqueryAPIRequest) GetQuery() *TpFundsSendQuery {
	return r._query
}
