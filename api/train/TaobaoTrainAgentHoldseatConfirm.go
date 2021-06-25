package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
火车票代理商接口——确认占座是否成功 
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功
*/
func TaobaoTrainAgentHoldseatConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentHoldseatConfirmRequest, session string) (*train.TaobaoTrainAgentHoldseatConfirmResponse, error) {
    var resp train.TaobaoTrainAgentHoldseatConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
