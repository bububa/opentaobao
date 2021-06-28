package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
线下票回填物流信息v2--增加鉴权校验 
taobao.train.agent.express.set.vtwo

线下票回填物流信息服务
*/
func TaobaoTrainAgentExpressSetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentExpressSetVtwoRequest, session string) (*train.TaobaoTrainAgentExpressSetVtwoAPIResponse, error) {
    var resp train.TaobaoTrainAgentExpressSetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
