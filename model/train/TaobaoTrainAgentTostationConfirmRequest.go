package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票确认送票至车站服务 APIRequest
taobao.train.agent.tostation.confirm

送票至车站的订单，代理商确认配送到站
*/
type TaobaoTrainAgentTostationConfirmRequest struct {
    model.Params

    // 淘宝的主订单号
    mainOrderId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentTostationConfirmRequest() *TaobaoTrainAgentTostationConfirmRequest{
    return &TaobaoTrainAgentTostationConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentTostationConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.tostation.confirm"
}

func (r TaobaoTrainAgentTostationConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentTostationConfirmRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentTostationConfirmRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentTostationConfirmRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentTostationConfirmRequest) GetAgentId() int64 {
    return r.agentId
}

