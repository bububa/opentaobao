package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundnewgetdetailAPIRequest 查询申请单详情(新版) API请求
// taobao.alitrip.ie.agent.refund.new.getdetail
//
// 查询申请单详情
type TaobaoalitripieagentrefundnewgetdetailAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq
}

// NewTaobaoalitripieagentrefundnewgetdetailRequest 初始化TaobaoalitripieagentrefundnewgetdetailAPIRequest对象
func NewTaobaoalitripieagentrefundnewgetdetailRequest() *TaobaoalitripieagentrefundnewgetdetailAPIRequest {
	return &TaobaoalitripieagentrefundnewgetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundnewgetdetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundnewgetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundnewgetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRefundOrderQueryDetailRq is ParamRefundOrderQueryDetailRq Setter
// 请求
func (r *TaobaoalitripieagentrefundnewgetdetailAPIRequest) SetParamRefundOrderQueryDetailRq(_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq) error {
	r._paramRefundOrderQueryDetailRq = _paramRefundOrderQueryDetailRq
	r.Set("param_refund_order_query_detail_rq", _paramRefundOrderQueryDetailRq)
	return nil
}

// GetParamRefundOrderQueryDetailRq ParamRefundOrderQueryDetailRq Getter
func (r TaobaoalitripieagentrefundnewgetdetailAPIRequest) GetParamRefundOrderQueryDetailRq() *RefundOrderQueryDetailRq {
	return r._paramRefundOrderQueryDetailRq
}
