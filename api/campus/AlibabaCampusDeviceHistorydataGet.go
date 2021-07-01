package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
设备历史数据批量获取 
alibaba.campus.device.historydata.get

设备历史数据批量获取
*/
func AlibabaCampusDeviceHistorydataGet(clt *core.SDKClient, req *campus.AlibabaCampusDeviceHistorydataGetAPIRequest, session string) (*campus.AlibabaCampusDeviceHistorydataGetAPIResponse, error) {
    var resp campus.AlibabaCampusDeviceHistorydataGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
