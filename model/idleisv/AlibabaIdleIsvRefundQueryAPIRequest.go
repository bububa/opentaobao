package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvrefundqueryAPIRequest 闲鱼已验货交易订单退款信息查询 API请求
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
type AlibabaidleisvrefundqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery
}

// NewAlibabaidleisvrefundqueryRequest 初始化AlibabaidleisvrefundqueryAPIRequest对象
func NewAlibabaidleisvrefundqueryRequest() *AlibabaidleisvrefundqueryAPIRequest {
	return &AlibabaidleisvrefundqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvrefundqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvrefundqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvrefundqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAppraiseIsvOrderQuery is ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaidleisvrefundqueryAPIRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
	r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
	r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
	return nil
}

// GetParamAppraiseIsvOrderQuery ParamAppraiseIsvOrderQuery Getter
func (r AlibabaidleisvrefundqueryAPIRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
	return r._paramAppraiseIsvOrderQuery
}
