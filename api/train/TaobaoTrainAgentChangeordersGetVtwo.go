package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
获取待改签订单v2--增加鉴权校验 
taobao.train.agent.changeorders.get.vtwo

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
func TaobaoTrainAgentChangeordersGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeordersGetVtwoRequest, session string) (*train.TaobaoTrainAgentChangeordersGetVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentChangeordersGetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
