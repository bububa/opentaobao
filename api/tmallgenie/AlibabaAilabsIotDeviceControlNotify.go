package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵IoT异步控制回调接口 
alibaba.ailabs.iot.device.control.notify

用于天猫精灵IoT云云接入控制结果的异步回调
*/
func AlibabaAilabsIotDeviceControlNotify(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceControlNotifyAPIRequest, session string) (*tmallgenie.AlibabaAilabsIotDeviceControlNotifyAPIResponse, error) {
    var resp tmallgenie.AlibabaAilabsIotDeviceControlNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
