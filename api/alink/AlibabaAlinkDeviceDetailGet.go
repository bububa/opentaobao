package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
获取设备详情 
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
func AlibabaAlinkDeviceDetailGet(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceDetailGetRequest, session string) (*alink.AlibabaAlinkDeviceDetailGetResponse, error) {
    var resp alink.AlibabaAlinkDeviceDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
