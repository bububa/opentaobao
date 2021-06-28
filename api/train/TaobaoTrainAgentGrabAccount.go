package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
代购抢代理商回传12306账号 
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
func TaobaoTrainAgentGrabAccount(clt *core.SDKClient, req *train.TaobaoTrainAgentGrabAccountRequest, session string) (*train.TaobaoTrainAgentGrabAccountAPIResponse, error) {
    var resp train.TaobaoTrainAgentGrabAccountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
