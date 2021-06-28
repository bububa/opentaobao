package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
获取mqtt设备信息 
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息
*/
func TaobaoWdkIotDeviceadminMqttDeviceGetwithtoken(clt *core.SDKClient, req *wdk.TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest, session string) (*wdk.TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse, error) {
    var resp wdk.TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
