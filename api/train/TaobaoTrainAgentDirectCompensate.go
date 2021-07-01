package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
func TaobaoTrainAgentDirectCompensate(clt *core.SDKClient, req *train.TaobaoTrainAgentDirectCompensateAPIRequest, session string) (*train.TaobaoTrainAgentDirectCompensateAPIResponse, error) {
    var resp train.TaobaoTrainAgentDirectCompensateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
