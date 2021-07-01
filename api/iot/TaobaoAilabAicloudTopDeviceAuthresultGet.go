package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取设备授权码验证结果 
taobao.ailab.aicloud.top.device.authresult.get

获取设备授权码验证结果
*/
func TaobaoAilabAicloudTopDeviceAuthresultGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
