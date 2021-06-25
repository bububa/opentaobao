package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代理商手动退款接口 
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口
*/
func TaobaoTrainAgentHandrefundRefundfee(clt *core.SDKClient, req *train.TaobaoTrainAgentHandrefundRefundfeeRequest, session string) (*train.TaobaoTrainAgentHandrefundRefundfeeResponse, error) {
    var resp train.TaobaoTrainAgentHandrefundRefundfeeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
