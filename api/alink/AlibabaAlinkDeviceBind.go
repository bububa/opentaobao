package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
绑定设备 
alibaba.alink.device.bind

阿里智能解绑设备
*/
func AlibabaAlinkDeviceBind(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceBindRequest, session string) (*alink.AlibabaAlinkDeviceBindAPIResponse, error) {
    var resp alink.AlibabaAlinkDeviceBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
