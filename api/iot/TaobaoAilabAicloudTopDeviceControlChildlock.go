package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
设备儿童锁 
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁
*/
func TaobaoAilabAicloudTopDeviceControlChildlock(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlChildlockRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlChildlockResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
