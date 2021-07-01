package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
toast 
alibaba.interact.sensor.toast

toast提示
*/
func AlibabaInteractSensorToast(clt *core.SDKClient, req *interact.AlibabaInteractSensorToastAPIRequest, session string) (*interact.AlibabaInteractSensorToastAPIResponse, error) {
    var resp interact.AlibabaInteractSensorToastAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
