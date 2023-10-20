package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlLampAPIResponse 台灯控制 API返回值
// taobao.ailab.aicloud.top.device.control.lamp
//
// 台灯控制
type TaobaoAilabAicloudTopDeviceControlLampAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlLampAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlLampAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlLampAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlLampAPIResponseModel is 台灯控制 成功返回结果
type TaobaoAilabAicloudTopDeviceControlLampAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_lamp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlLampAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceControlLampAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlLampAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlLampAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlLampAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlLampAPIResponse() *TaobaoAilabAicloudTopDeviceControlLampAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlLampAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlLampAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlLampAPIResponse 将 TaobaoAilabAicloudTopDeviceControlLampAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlLampAPIResponse(v *TaobaoAilabAicloudTopDeviceControlLampAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlLampAPIResponse.Put(v)
}
