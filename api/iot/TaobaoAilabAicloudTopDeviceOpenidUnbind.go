package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
openTaoBaoId解绑设备 
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
func TaobaoAilabAicloudTopDeviceOpenidUnbind(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
