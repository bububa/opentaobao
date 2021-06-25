package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
码相关API 
alibaba.interact.sensor.ma

码相关API
*/
func AlibabaInteractSensorMa(clt *core.SDKClient, req *interact.AlibabaInteractSensorMaRequest, session string) (*interact.AlibabaInteractSensorMaResponse, error) {
    var resp interact.AlibabaInteractSensorMaAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
