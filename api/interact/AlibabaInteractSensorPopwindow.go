package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
popwindow 
alibaba.interact.sensor.popwindow

popwindow
*/
func AlibabaInteractSensorPopwindow(clt *core.SDKClient, req *interact.AlibabaInteractSensorPopwindowRequest, session string) (*interact.AlibabaInteractSensorPopwindowResponse, error) {
    var resp interact.AlibabaInteractSensorPopwindowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
