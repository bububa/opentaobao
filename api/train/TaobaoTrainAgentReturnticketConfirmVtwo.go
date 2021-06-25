package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
退票通知 
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
func TaobaoTrainAgentReturnticketConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentReturnticketConfirmVtwoRequest, session string) (*train.TaobaoTrainAgentReturnticketConfirmVtwoResponse, error) {
    var resp train.TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
