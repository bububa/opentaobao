package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse 开放设备id转换内部设备id API返回值
// taobao.ailab.aicloud.top.device.deviceid.convert
//
// 将开放设备id转换为内部设备id
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponseModel is 开放设备id转换内部设备id 成功返回结果
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_deviceid_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAilabAicloudTopDeviceDeviceidConvertResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse
func GetTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse() *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse 将 TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse(v *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse.Put(v)
}
