package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
重力感应 
alibaba.interact.sensor.gravity

native获取重力感应
*/
func AlibabaInteractSensorGravity(clt *core.SDKClient, req *interact.AlibabaInteractSensorGravityRequest, session string) (*interact.AlibabaInteractSensorGravityAPIResponse, error) {
    var resp interact.AlibabaInteractSensorGravityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
