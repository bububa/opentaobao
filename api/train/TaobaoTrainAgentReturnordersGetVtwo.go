package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
获取待退票的订单v2--增加鉴权校验 
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。
*/
func TaobaoTrainAgentReturnordersGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentReturnordersGetVtwoRequest, session string) (*train.TaobaoTrainAgentReturnordersGetVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentReturnordersGetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
