package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代理商获取订单信息回调APIv2--增加鉴权校验 
taobao.train.agent.order.get.vtwo

代理商获取订单信息回调API
*/
func TaobaoTrainAgentOrderGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderGetVtwoRequest, session string) (*train.TaobaoTrainAgentOrderGetVtwoResponse, error) {
    var resp train.TaobaoTrainAgentOrderGetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
