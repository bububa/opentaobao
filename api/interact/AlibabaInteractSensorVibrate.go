package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
客户端震动 
alibaba.interact.sensor.vibrate

客户端震动
*/
func AlibabaInteractSensorVibrate(clt *core.SDKClient, req *interact.AlibabaInteractSensorVibrateRequest, session string) (*interact.AlibabaInteractSensorVibrateResponse, error) {
    var resp interact.AlibabaInteractSensorVibrateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
