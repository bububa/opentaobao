package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
点播url 
taobao.ailab.aicloud.top.device.control.playurl

点播url
*/
func TaobaoAilabAicloudTopDeviceControlPlayurl(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPlayurlRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
