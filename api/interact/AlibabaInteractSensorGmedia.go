package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
gmedia 
alibaba.interact.sensor.gmedia

媒体功能
*/
func AlibabaInteractSensorGmedia(clt *core.SDKClient, req *interact.AlibabaInteractSensorGmediaAPIRequest, session string) (*interact.AlibabaInteractSensorGmediaAPIResponse, error) {
    var resp interact.AlibabaInteractSensorGmediaAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
