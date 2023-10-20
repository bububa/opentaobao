package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceStatusUpdateAPIResponse ailabs iot 设备状态更新 API返回值
// alibaba.ailabs.iot.device.status.update
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
type AlibabaAilabsIotDeviceStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel).Reset()
}

// AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel is ailabs iot 设备状态更新 成功返回结果
type AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备状态更新是否成功
	Result *AlibabaAilabsIotDeviceStatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsIotDeviceStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceStatusUpdateAPIResponse)
	},
}

// GetAlibabaAilabsIotDeviceStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotDeviceStatusUpdateAPIResponse
func GetAlibabaAilabsIotDeviceStatusUpdateAPIResponse() *AlibabaAilabsIotDeviceStatusUpdateAPIResponse {
	return poolAlibabaAilabsIotDeviceStatusUpdateAPIResponse.Get().(*AlibabaAilabsIotDeviceStatusUpdateAPIResponse)
}

// ReleaseAlibabaAilabsIotDeviceStatusUpdateAPIResponse 将 AlibabaAilabsIotDeviceStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotDeviceStatusUpdateAPIResponse(v *AlibabaAilabsIotDeviceStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotDeviceStatusUpdateAPIResponse.Put(v)
}
