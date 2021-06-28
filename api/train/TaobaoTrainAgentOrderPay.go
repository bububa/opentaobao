package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代购订单代付接口 
taobao.train.agent.order.pay

代购订单代付接口
*/
func TaobaoTrainAgentOrderPay(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderPayRequest, session string) (*train.TaobaoTrainAgentOrderPayAPIResponse, error) {
    var resp train.TaobaoTrainAgentOrderPayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
