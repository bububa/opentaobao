package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
获取待改签订单 
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
func TaobaoTrainAgentChangeordersGet(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeordersGetAPIRequest, session string) (*train.TaobaoTrainAgentChangeordersGetAPIResponse, error) {
    var resp train.TaobaoTrainAgentChangeordersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
