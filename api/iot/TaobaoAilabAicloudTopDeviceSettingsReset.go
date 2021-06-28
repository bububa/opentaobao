package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
重置设备个性化设置 
taobao.ailab.aicloud.top.device.settings.reset

重置设备个性化设置
*/
func TaobaoAilabAicloudTopDeviceSettingsReset(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceSettingsResetRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
