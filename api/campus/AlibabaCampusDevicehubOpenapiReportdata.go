package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
设备数据上报 
alibaba.campus.devicehub.openapi.reportdata

设备数据上报
*/
func AlibabaCampusDevicehubOpenapiReportdata(clt *core.SDKClient, req *campus.AlibabaCampusDevicehubOpenapiReportdataRequest, session string) (*campus.AlibabaCampusDevicehubOpenapiReportdataAPIResponse, error) {
    var resp campus.AlibabaCampusDevicehubOpenapiReportdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
