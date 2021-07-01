package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
获取改签单详情v2--增加鉴权校验 
taobao.train.agent.change.get.vtwo

卖家获取待处理的改签单详情
*/
func TaobaoTrainAgentChangeGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeGetVtwoAPIRequest, session string) (*train.TaobaoTrainAgentChangeGetVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentChangeGetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
