package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse 设备列表更新通知 API返回值
// alibaba.ailabs.iot.device.list.update.notify
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
type AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel).Reset()
}

// AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel is 设备列表更新通知 成功返回结果
type AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_list_update_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsIotDeviceListUpdateNotifyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse)
	},
}

// GetAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse
func GetAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse() *AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse {
	return poolAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse.Get().(*AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse)
}

// ReleaseAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse 将 AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse(v *AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotDeviceListUpdateNotifyAPIResponse.Put(v)
}
