package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
多条件查询设备分组 
alibaba.campus.device.openapi.getdevicelist

多条件查询设备分组
*/
func AlibabaCampusDeviceOpenapiGetdevicelist(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
