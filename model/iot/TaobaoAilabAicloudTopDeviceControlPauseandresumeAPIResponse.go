package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse 设备播放暂停 API返回值
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel is 设备播放暂停 成功返回结果
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_pauseandresume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse() *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse 将 TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse(v *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse.Put(v)
}
