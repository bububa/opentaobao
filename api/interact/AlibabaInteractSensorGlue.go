package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
视频播放 
alibaba.interact.sensor.glue

视频播放
*/
func AlibabaInteractSensorGlue(clt *core.SDKClient, req *interact.AlibabaInteractSensorGlueRequest, session string) (*interact.AlibabaInteractSensorGlueResponse, error) {
    var resp interact.AlibabaInteractSensorGlueAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
