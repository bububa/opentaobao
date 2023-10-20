package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundnewgetlistAPIRequest 新查询退票申请单列表 API请求
// taobao.alitrip.ie.agent.refund.new.getlist
//
// 查询商家国际机票退票申请单列表
type TaobaoalitripieagentrefundnewgetlistAPIRequest struct {
	model.Params
	// 列表查询
	_paramRefundOrderQueryListRq *RefundOrderQueryListRq
}

// NewTaobaoalitripieagentrefundnewgetlistRequest 初始化TaobaoalitripieagentrefundnewgetlistAPIRequest对象
func NewTaobaoalitripieagentrefundnewgetlistRequest() *TaobaoalitripieagentrefundnewgetlistAPIRequest {
	return &TaobaoalitripieagentrefundnewgetlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundnewgetlistAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundnewgetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundnewgetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRefundOrderQueryListRq is ParamRefundOrderQueryListRq Setter
// 列表查询
func (r *TaobaoalitripieagentrefundnewgetlistAPIRequest) SetParamRefundOrderQueryListRq(_paramRefundOrderQueryListRq *RefundOrderQueryListRq) error {
	r._paramRefundOrderQueryListRq = _paramRefundOrderQueryListRq
	r.Set("param_refund_order_query_list_rq", _paramRefundOrderQueryListRq)
	return nil
}

// GetParamRefundOrderQueryListRq ParamRefundOrderQueryListRq Getter
func (r TaobaoalitripieagentrefundnewgetlistAPIRequest) GetParamRefundOrderQueryListRq() *RefundOrderQueryListRq {
	return r._paramRefundOrderQueryListRq
}
