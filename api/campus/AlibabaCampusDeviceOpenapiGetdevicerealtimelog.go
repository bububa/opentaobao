package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据设备uuid获取设备采集信息 
alibaba.campus.device.openapi.getdevicerealtimelog

根据设备uuid获取设备采集信息
*/
func AlibabaCampusDeviceOpenapiGetdevicerealtimelog(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
