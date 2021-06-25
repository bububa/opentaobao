package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单信息回调API APIRequest
taobao.train.agent.order.get

代理商获取订单信息回调API
*/
type TaobaoTrainAgentOrderGetRequest struct {
    model.Params

    // 淘宝的主订单号
    mainOrderId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentOrderGetRequest() *TaobaoTrainAgentOrderGetRequest{
    return &TaobaoTrainAgentOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentOrderGetRequest) GetApiMethodName() string {
    return "taobao.train.agent.order.get"
}

func (r TaobaoTrainAgentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentOrderGetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentOrderGetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentOrderGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentOrderGetRequest) GetAgentId() int64 {
    return r.agentId
}

