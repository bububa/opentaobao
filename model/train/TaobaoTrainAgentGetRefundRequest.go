package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单退票信息 APIRequest
taobao.train.agent.get.refund

代理商获取订单信息回调API
*/
type TaobaoTrainAgentGetRefundRequest struct {
    model.Params

    // 淘宝的主订单号
    mainOrderId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentGetRefundRequest() *TaobaoTrainAgentGetRefundRequest{
    return &TaobaoTrainAgentGetRefundRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentGetRefundRequest) GetApiMethodName() string {
    return "taobao.train.agent.get.refund"
}

func (r TaobaoTrainAgentGetRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentGetRefundRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentGetRefundRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentGetRefundRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentGetRefundRequest) GetAgentId() int64 {
    return r.agentId
}

