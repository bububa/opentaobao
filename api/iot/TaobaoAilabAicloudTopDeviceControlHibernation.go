package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
定时休眠 
taobao.ailab.aicloud.top.device.control.hibernation

定时休眠
*/
func TaobaoAilabAicloudTopDeviceControlHibernation(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlHibernationRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
