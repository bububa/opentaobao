package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensoraudioAPIResponse 声音 API返回值
// alibaba.interact.sensor.audio
//
// 客户端声音
type AlibabainteractsensoraudioAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensoraudioAPIResponseModel
}

// AlibabainteractsensoraudioAPIResponseModel is 声音 成功返回结果
type AlibabainteractsensoraudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_audio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
