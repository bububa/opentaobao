package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
网络状态 
alibaba.interact.sensor.networkstatus

客户端网络状态
*/
func AlibabaInteractSensorNetworkstatus(clt *core.SDKClient, req *interact.AlibabaInteractSensorNetworkstatusRequest, session string) (*interact.AlibabaInteractSensorNetworkstatusResponse, error) {
    var resp interact.AlibabaInteractSensorNetworkstatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
