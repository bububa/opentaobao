package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceGetstatusAPIResponse 获取设备状态 API返回值
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
type TaobaoAilabAicloudTopDeviceGetstatusAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceGetstatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceGetstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceGetstatusAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceGetstatusAPIResponseModel is 获取设备状态 成功返回结果
type TaobaoAilabAicloudTopDeviceGetstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_getstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceGetstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceGetstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceGetstatusAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceGetstatusAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceGetstatusAPIResponse
func GetTaobaoAilabAicloudTopDeviceGetstatusAPIResponse() *TaobaoAilabAicloudTopDeviceGetstatusAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceGetstatusAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceGetstatusAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceGetstatusAPIResponse 将 TaobaoAilabAicloudTopDeviceGetstatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceGetstatusAPIResponse(v *TaobaoAilabAicloudTopDeviceGetstatusAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceGetstatusAPIResponse.Put(v)
}
