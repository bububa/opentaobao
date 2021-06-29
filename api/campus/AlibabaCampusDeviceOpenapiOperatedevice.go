package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据uuid操作设备 
alibaba.campus.device.openapi.operatedevice

根据uuid操作设备
*/
func AlibabaCampusDeviceOpenapiOperatedevice(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiOperatedeviceRequest, session string) (*campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
