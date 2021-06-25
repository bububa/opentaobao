package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
台灯控制 
taobao.ailab.aicloud.top.device.control.lamp

台灯控制
*/
func TaobaoAilabAicloudTopDeviceControlLamp(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlLampRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlLampResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlLampAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
