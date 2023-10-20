package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceControlNotifyAPIResponse 天猫精灵IoT异步控制回调接口 API返回值
// alibaba.ailabs.iot.device.control.notify
//
// 用于天猫精灵IoT云云接入控制结果的异步回调
type AlibabaAilabsIotDeviceControlNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceControlNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceControlNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotDeviceControlNotifyAPIResponseModel).Reset()
}

// AlibabaAilabsIotDeviceControlNotifyAPIResponseModel is 天猫精灵IoT异步控制回调接口 成功返回结果
type AlibabaAilabsIotDeviceControlNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_control_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否通知成功
	RetValue bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceControlNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetValue = false
}

var poolAlibabaAilabsIotDeviceControlNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceControlNotifyAPIResponse)
	},
}

// GetAlibabaAilabsIotDeviceControlNotifyAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotDeviceControlNotifyAPIResponse
func GetAlibabaAilabsIotDeviceControlNotifyAPIResponse() *AlibabaAilabsIotDeviceControlNotifyAPIResponse {
	return poolAlibabaAilabsIotDeviceControlNotifyAPIResponse.Get().(*AlibabaAilabsIotDeviceControlNotifyAPIResponse)
}

// ReleaseAlibabaAilabsIotDeviceControlNotifyAPIResponse 将 AlibabaAilabsIotDeviceControlNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotDeviceControlNotifyAPIResponse(v *AlibabaAilabsIotDeviceControlNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotDeviceControlNotifyAPIResponse.Put(v)
}
