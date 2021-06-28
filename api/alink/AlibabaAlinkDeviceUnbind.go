package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
解绑设备 
alibaba.alink.device.unbind

阿里智能解绑设备
*/
func AlibabaAlinkDeviceUnbind(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnbindRequest, session string) (*alink.AlibabaAlinkDeviceUnbindAPIResponse, error) {
    var resp alink.AlibabaAlinkDeviceUnbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
