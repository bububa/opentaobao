package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListGetAPIResponse 获取iot设备列表 API返回值
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
type AlibabaAilabsIotDeviceListGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotDeviceListGetAPIResponseModel).Reset()
}

// AlibabaAilabsIotDeviceListGetAPIResponseModel is 获取iot设备列表 成功返回结果
type AlibabaAilabsIotDeviceListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsIotDeviceListGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsIotDeviceListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceListGetAPIResponse)
	},
}

// GetAlibabaAilabsIotDeviceListGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotDeviceListGetAPIResponse
func GetAlibabaAilabsIotDeviceListGetAPIResponse() *AlibabaAilabsIotDeviceListGetAPIResponse {
	return poolAlibabaAilabsIotDeviceListGetAPIResponse.Get().(*AlibabaAilabsIotDeviceListGetAPIResponse)
}

// ReleaseAlibabaAilabsIotDeviceListGetAPIResponse 将 AlibabaAilabsIotDeviceListGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotDeviceListGetAPIResponse(v *AlibabaAilabsIotDeviceListGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotDeviceListGetAPIResponse.Put(v)
}
