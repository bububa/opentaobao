package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
隐藏titleBar 
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
func AlibabaInteractSensorTitlebarhide(clt *core.SDKClient, req *interact.AlibabaInteractSensorTitlebarhideRequest, session string) (*interact.AlibabaInteractSensorTitlebarhideResponse, error) {
    var resp interact.AlibabaInteractSensorTitlebarhideAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
