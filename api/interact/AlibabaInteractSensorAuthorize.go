package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
客户端授权页 
alibaba.interact.sensor.authorize

客户端授权页
*/
func AlibabaInteractSensorAuthorize(clt *core.SDKClient, req *interact.AlibabaInteractSensorAuthorizeRequest, session string) (*interact.AlibabaInteractSensorAuthorizeAPIResponse, error) {
    var resp interact.AlibabaInteractSensorAuthorizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
