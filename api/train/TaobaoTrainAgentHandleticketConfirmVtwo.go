package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代理商出票中v2--增加鉴权校验 
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中
*/
func TaobaoTrainAgentHandleticketConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentHandleticketConfirmVtwoRequest, session string) (*train.TaobaoTrainAgentHandleticketConfirmVtwoResponse, error) {
    var resp train.TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
