package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询申请单详情(新版) API请求
taobao.alitrip.ie.agent.refund.new.getdetail

查询申请单详情
*/
type TaobaoAlitripIeAgentRefundNewGetdetailRequest struct {
    model.Params
    // 请求
    paramRefundOrderQueryDetailRq   *RefundOrderQueryDetailRq
}

// 初始化TaobaoAlitripIeAgentRefundNewGetdetailRequest对象
func NewTaobaoAlitripIeAgentRefundNewGetdetailRequest() *TaobaoAlitripIeAgentRefundNewGetdetailRequest{
    return &TaobaoAlitripIeAgentRefundNewGetdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewGetdetailRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRefundOrderQueryDetailRq Setter
// 请求
func (r *TaobaoAlitripIeAgentRefundNewGetdetailRequest) SetParamRefundOrderQueryDetailRq(paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq) error {
    r.paramRefundOrderQueryDetailRq = paramRefundOrderQueryDetailRq
    r.Set("param_refund_order_query_detail_rq", paramRefundOrderQueryDetailRq)
    return nil
}

// ParamRefundOrderQueryDetailRq Getter
func (r TaobaoAlitripIeAgentRefundNewGetdetailRequest) GetParamRefundOrderQueryDetailRq() *RefundOrderQueryDetailRq {
    return r.paramRefundOrderQueryDetailRq
}
