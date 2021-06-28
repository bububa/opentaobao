package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
分享 
alibaba.interact.sensor.share

客户端分享
*/
func AlibabaInteractSensorShare(clt *core.SDKClient, req *interact.AlibabaInteractSensorShareRequest, session string) (*interact.AlibabaInteractSensorShareAPIResponse, error) {
    var resp interact.AlibabaInteractSensorShareAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
