package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取设备状态信息 
taobao.ailab.aicloud.top.device.statusinfo.get

获取设备状态信息
*/
func TaobaoAilabAicloudTopDeviceStatusinfoGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceStatusinfoGetRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceStatusinfoGetResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
