package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取设备授权码 
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
func TaobaoAilabAicloudTopDeviceAuthcodeGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
