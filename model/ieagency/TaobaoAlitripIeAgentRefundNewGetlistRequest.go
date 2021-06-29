package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新查询退票申请单列表 API请求
taobao.alitrip.ie.agent.refund.new.getlist

查询商家国际机票退票申请单列表
*/
type TaobaoAlitripIeAgentRefundNewGetlistRequest struct {
    model.Params
    // 列表查询
    paramRefundOrderQueryListRq   *RefundOrderQueryListRq
}

// 初始化TaobaoAlitripIeAgentRefundNewGetlistRequest对象
func NewTaobaoAlitripIeAgentRefundNewGetlistRequest() *TaobaoAlitripIeAgentRefundNewGetlistRequest{
    return &TaobaoAlitripIeAgentRefundNewGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewGetlistRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.getlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRefundOrderQueryListRq Setter
// 列表查询
func (r *TaobaoAlitripIeAgentRefundNewGetlistRequest) SetParamRefundOrderQueryListRq(paramRefundOrderQueryListRq *RefundOrderQueryListRq) error {
    r.paramRefundOrderQueryListRq = paramRefundOrderQueryListRq
    r.Set("param_refund_order_query_list_rq", paramRefundOrderQueryListRq)
    return nil
}

// ParamRefundOrderQueryListRq Getter
func (r TaobaoAlitripIeAgentRefundNewGetlistRequest) GetParamRefundOrderQueryListRq() *RefundOrderQueryListRq {
    return r.paramRefundOrderQueryListRq
}
