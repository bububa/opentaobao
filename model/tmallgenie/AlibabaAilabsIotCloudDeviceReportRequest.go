package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵云云接入设备状态、事件上报接口 API请求
alibaba.ailabs.iot.cloud.device.report

承接天猫精灵云云接入设备的状态、事件上报
*/
type AlibabaAilabsIotCloudDeviceReportRequest struct {
    model.Params
    // 上报总入参
    _cloudReportParam   *CloudReportParam
}

// 初始化AlibabaAilabsIotCloudDeviceReportRequest对象
func NewAlibabaAilabsIotCloudDeviceReportRequest() *AlibabaAilabsIotCloudDeviceReportRequest{
    return &AlibabaAilabsIotCloudDeviceReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotCloudDeviceReportRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.cloud.device.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotCloudDeviceReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CloudReportParam Setter
// 上报总入参
func (r *AlibabaAilabsIotCloudDeviceReportRequest) SetCloudReportParam(_cloudReportParam *CloudReportParam) error {
    r._cloudReportParam = _cloudReportParam
    r.Set("cloud_report_param", _cloudReportParam)
    return nil
}

// CloudReportParam Getter
func (r AlibabaAilabsIotCloudDeviceReportRequest) GetCloudReportParam() *CloudReportParam {
    return r._cloudReportParam
}
