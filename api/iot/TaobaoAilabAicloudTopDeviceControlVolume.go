package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
设备音量 
taobao.ailab.aicloud.top.device.control.volume

设备音量
*/
func TaobaoAilabAicloudTopDeviceControlVolume(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlVolumeRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlVolumeResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
