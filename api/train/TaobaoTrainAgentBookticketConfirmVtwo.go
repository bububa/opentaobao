package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
火车票代理商接口——确认出票是否成功v2--增加鉴权校验 
taobao.train.agent.bookticket.confirm.vtwo

火车票代理商接口——确认出票是否成功
*/
func TaobaoTrainAgentBookticketConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentBookticketConfirmVtwoRequest, session string) (*train.TaobaoTrainAgentBookticketConfirmVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentBookticketConfirmVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
