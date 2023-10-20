package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse 通过id播放歌曲 API返回值
// taobao.ailab.aicloud.top.device.control.playbyid
//
// 通过id播放歌曲
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponseModel is 通过id播放歌曲 成功返回结果
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_playbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse() *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse 将 TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse(v *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse.Put(v)
}
