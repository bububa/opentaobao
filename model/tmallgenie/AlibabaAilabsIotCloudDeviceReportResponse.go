package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵云云接入设备状态、事件上报接口 APIResponse
alibaba.ailabs.iot.cloud.device.report

承接天猫精灵云云接入设备的状态、事件上报
*/
type AlibabaAilabsIotCloudDeviceReportAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotCloudDeviceReportResponse
}

type AlibabaAilabsIotCloudDeviceReportResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_cloud_device_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值，是否上报成功
    
    RetValue   bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`

    
}
