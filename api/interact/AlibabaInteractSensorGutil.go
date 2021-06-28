package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
canvas工具包 
alibaba.interact.sensor.gutil

canvas工具包
*/
func AlibabaInteractSensorGutil(clt *core.SDKClient, req *interact.AlibabaInteractSensorGutilRequest, session string) (*interact.AlibabaInteractSensorGutilAPIResponse, error) {
    var resp interact.AlibabaInteractSensorGutilAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
