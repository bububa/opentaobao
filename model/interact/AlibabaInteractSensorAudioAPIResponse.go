package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorAudioAPIResponse 声音 API返回值
// alibaba.interact.sensor.audio
//
// 客户端声音
type AlibabaInteractSensorAudioAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorAudioAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorAudioAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorAudioAPIResponseModel).Reset()
}

// AlibabaInteractSensorAudioAPIResponseModel is 声音 成功返回结果
type AlibabaInteractSensorAudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_audio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorAudioAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorAudioAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorAudioAPIResponse)
	},
}

// GetAlibabaInteractSensorAudioAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorAudioAPIResponse
func GetAlibabaInteractSensorAudioAPIResponse() *AlibabaInteractSensorAudioAPIResponse {
	return poolAlibabaInteractSensorAudioAPIResponse.Get().(*AlibabaInteractSensorAudioAPIResponse)
}

// ReleaseAlibabaInteractSensorAudioAPIResponse 将 AlibabaInteractSensorAudioAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorAudioAPIResponse(v *AlibabaInteractSensorAudioAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorAudioAPIResponse.Put(v)
}
