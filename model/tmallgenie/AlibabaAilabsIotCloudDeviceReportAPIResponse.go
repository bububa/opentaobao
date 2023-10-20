package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotclouddevicereportAPIResponse 天猫精灵云云接入设备状态、事件上报接口 API返回值
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
type AlibabaailabsiotclouddevicereportAPIResponse struct {
	model.CommonResponse
	AlibabaailabsiotclouddevicereportAPIResponseModel
}

// AlibabaailabsiotclouddevicereportAPIResponseModel is 天猫精灵云云接入设备状态、事件上报接口 成功返回结果
type AlibabaailabsiotclouddevicereportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_cloud_device_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值，是否上报成功
	RetValue bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}
