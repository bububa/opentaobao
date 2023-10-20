package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractmediaaudioAPIResponse 音频相关鉴权接口 API返回值
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
type AlibabainteractmediaaudioAPIResponse struct {
	model.CommonResponse
	AlibabainteractmediaaudioAPIResponseModel
}

// AlibabainteractmediaaudioAPIResponseModel is 音频相关鉴权接口 成功返回结果
type AlibabainteractmediaaudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_media_audio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
