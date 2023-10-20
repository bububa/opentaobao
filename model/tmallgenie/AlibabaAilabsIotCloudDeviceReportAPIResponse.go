package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotCloudDeviceReportAPIResponse 天猫精灵云云接入设备状态、事件上报接口 API返回值
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
type AlibabaAilabsIotCloudDeviceReportAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotCloudDeviceReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotCloudDeviceReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotCloudDeviceReportAPIResponseModel).Reset()
}

// AlibabaAilabsIotCloudDeviceReportAPIResponseModel is 天猫精灵云云接入设备状态、事件上报接口 成功返回结果
type AlibabaAilabsIotCloudDeviceReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_cloud_device_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值，是否上报成功
	RetValue bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotCloudDeviceReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetValue = false
}

var poolAlibabaAilabsIotCloudDeviceReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotCloudDeviceReportAPIResponse)
	},
}

// GetAlibabaAilabsIotCloudDeviceReportAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotCloudDeviceReportAPIResponse
func GetAlibabaAilabsIotCloudDeviceReportAPIResponse() *AlibabaAilabsIotCloudDeviceReportAPIResponse {
	return poolAlibabaAilabsIotCloudDeviceReportAPIResponse.Get().(*AlibabaAilabsIotCloudDeviceReportAPIResponse)
}

// ReleaseAlibabaAilabsIotCloudDeviceReportAPIResponse 将 AlibabaAilabsIotCloudDeviceReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotCloudDeviceReportAPIResponse(v *AlibabaAilabsIotCloudDeviceReportAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotCloudDeviceReportAPIResponse.Put(v)
}
