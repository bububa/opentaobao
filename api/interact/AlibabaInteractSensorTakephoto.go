package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
takePhoto 
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
func AlibabaInteractSensorTakephoto(clt *core.SDKClient, req *interact.AlibabaInteractSensorTakephotoRequest, session string) (*interact.AlibabaInteractSensorTakephotoResponse, error) {
    var resp interact.AlibabaInteractSensorTakephotoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
