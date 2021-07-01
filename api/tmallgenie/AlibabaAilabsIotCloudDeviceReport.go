package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵云云接入设备状态、事件上报接口 
alibaba.ailabs.iot.cloud.device.report

承接天猫精灵云云接入设备的状态、事件上报
*/
func AlibabaAilabsIotCloudDeviceReport(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotCloudDeviceReportAPIRequest, session string) (*tmallgenie.AlibabaAilabsIotCloudDeviceReportAPIResponse, error) {
    var resp tmallgenie.AlibabaAilabsIotCloudDeviceReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
