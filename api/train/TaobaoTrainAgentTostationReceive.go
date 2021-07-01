package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
线下票送票至车站代理商确认用户已取票服务 
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
func TaobaoTrainAgentTostationReceive(clt *core.SDKClient, req *train.TaobaoTrainAgentTostationReceiveAPIRequest, session string) (*train.TaobaoTrainAgentTostationReceiveAPIResponse, error) {
    var resp train.TaobaoTrainAgentTostationReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
