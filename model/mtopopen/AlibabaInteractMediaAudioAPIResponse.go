package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractMediaAudioAPIResponse 音频相关鉴权接口 API返回值
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
type AlibabaInteractMediaAudioAPIResponse struct {
	model.CommonResponse
	AlibabaInteractMediaAudioAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractMediaAudioAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractMediaAudioAPIResponseModel).Reset()
}

// AlibabaInteractMediaAudioAPIResponseModel is 音频相关鉴权接口 成功返回结果
type AlibabaInteractMediaAudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_media_audio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractMediaAudioAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaInteractMediaAudioAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractMediaAudioAPIResponse)
	},
}

// GetAlibabaInteractMediaAudioAPIResponse 从 sync.Pool 获取 AlibabaInteractMediaAudioAPIResponse
func GetAlibabaInteractMediaAudioAPIResponse() *AlibabaInteractMediaAudioAPIResponse {
	return poolAlibabaInteractMediaAudioAPIResponse.Get().(*AlibabaInteractMediaAudioAPIResponse)
}

// ReleaseAlibabaInteractMediaAudioAPIResponse 将 AlibabaInteractMediaAudioAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractMediaAudioAPIResponse(v *AlibabaInteractMediaAudioAPIResponse) {
	v.Reset()
	poolAlibabaInteractMediaAudioAPIResponse.Put(v)
}
