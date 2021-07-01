package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询设备历史数据 
alibaba.campus.device.openapi.gethistorydata

查询历史数据的接口
*/
func AlibabaCampusDeviceOpenapiGethistorydata(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGethistorydataAPIRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGethistorydataAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceOpenapiGethistorydataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
