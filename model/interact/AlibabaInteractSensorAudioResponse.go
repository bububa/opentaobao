package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
声音 APIResponse
alibaba.interact.sensor.audio

客户端声音
*/
type AlibabaInteractSensorAudioAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorAudioResponse
}

type AlibabaInteractSensorAudioResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_audio_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
