package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵云云接入设备状态、事件上报接口 APIRequest
alibaba.ailabs.iot.cloud.device.report

承接天猫精灵云云接入设备的状态、事件上报
*/
type AlibabaAilabsIotCloudDeviceReportRequest struct {
    model.Params

    // 上报总入参
    cloudReportParam   *CloudReportParam 

}

func NewAlibabaAilabsIotCloudDeviceReportRequest() *AlibabaAilabsIotCloudDeviceReportRequest{
    return &AlibabaAilabsIotCloudDeviceReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotCloudDeviceReportRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.cloud.device.report"
}

func (r AlibabaAilabsIotCloudDeviceReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotCloudDeviceReportRequest) SetCloudReportParam(cloudReportParam *CloudReportParam) error {
    r.cloudReportParam = cloudReportParam
    r.Set("cloud_report_param", cloudReportParam)
    return nil
}

func (r AlibabaAilabsIotCloudDeviceReportRequest) GetCloudReportParam() *CloudReportParam {
    return r.cloudReportParam
}

