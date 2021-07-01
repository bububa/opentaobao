package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
火车票代理商接口——确认出票是否成功 
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功
*/
func TaobaoTrainAgentBookticketConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentBookticketConfirmAPIRequest, session string) (*train.TaobaoTrainAgentBookticketConfirmAPIResponse, error) {
    var resp train.TaobaoTrainAgentBookticketConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
