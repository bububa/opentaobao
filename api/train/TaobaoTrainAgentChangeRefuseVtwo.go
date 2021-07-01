package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代理商拒绝改签v2--增加鉴权校验 
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务
*/
func TaobaoTrainAgentChangeRefuseVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeRefuseVtwoAPIRequest, session string) (*train.TaobaoTrainAgentChangeRefuseVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentChangeRefuseVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
