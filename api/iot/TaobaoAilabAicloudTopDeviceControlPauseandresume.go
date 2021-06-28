package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
设备播放暂停 
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
func TaobaoAilabAicloudTopDeviceControlPauseandresume(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
