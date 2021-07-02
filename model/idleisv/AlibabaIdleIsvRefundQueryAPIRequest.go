package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvRefundQueryAPIRequest 闲鱼已验货交易订单退款信息查询 API请求
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
type AlibabaIdleIsvRefundQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery
}

// NewAlibabaIdleIsvRefundQueryRequest 初始化AlibabaIdleIsvRefundQueryAPIRequest对象
func NewAlibabaIdleIsvRefundQueryRequest() *AlibabaIdleIsvRefundQueryAPIRequest {
	return &AlibabaIdleIsvRefundQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamAppraiseIsvOrderQuery is ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaIdleIsvRefundQueryAPIRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
	r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
	r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
	return nil
}

// GetParamAppraiseIsvOrderQuery ParamAppraiseIsvOrderQuery Getter
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
	return r._paramAppraiseIsvOrderQuery
}
