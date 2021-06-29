package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
获取指定设备下指定参数的实时值 
alibaba.campus.device.openapi.getdevicerealtimedata

获取指定设备下指定参数的实时值
*/
func AlibabaCampusDeviceOpenapiGetdevicerealtimedata(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
