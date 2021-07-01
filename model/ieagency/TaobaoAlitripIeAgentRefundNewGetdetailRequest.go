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
type TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest struct {
    model.Params
    // 请求
    _paramRefundOrderQueryDetailRq   *RefundOrderQueryDetailRq
}

// 初始化TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest对象
func NewTaobaoAlitripIeAgentRefundNewGetdetailRequest() *TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest{
    return &TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRefundOrderQueryDetailRq Setter
// 请求
func (r *TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) SetParamRefundOrderQueryDetailRq(_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq) error {
    r._paramRefundOrderQueryDetailRq = _paramRefundOrderQueryDetailRq
    r.Set("param_refund_order_query_detail_rq", _paramRefundOrderQueryDetailRq)
    return nil
}

// ParamRefundOrderQueryDetailRq Getter
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetParamRefundOrderQueryDetailRq() *RefundOrderQueryDetailRq {
    return r._paramRefundOrderQueryDetailRq
}
