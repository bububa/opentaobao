package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票送票至车站代理商确认用户已取票服务 APIRequest
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
type TaobaoTrainAgentTostationReceiveRequest struct {
    model.Params

    // 淘宝的主订单号
    mainOrderId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentTostationReceiveRequest() *TaobaoTrainAgentTostationReceiveRequest{
    return &TaobaoTrainAgentTostationReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentTostationReceiveRequest) GetApiMethodName() string {
    return "taobao.train.agent.tostation.receive"
}

func (r TaobaoTrainAgentTostationReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentTostationReceiveRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentTostationReceiveRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentTostationReceiveRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentTostationReceiveRequest) GetAgentId() int64 {
    return r.agentId
}

