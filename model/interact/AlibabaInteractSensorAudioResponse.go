package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
声音 APIResponse
alibaba.interact.sensor.audio

客户端声音
*/
type AlibabaInteractSensorAudioAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorAudioResponse `json:"alibaba_interact_sensor_audio_response,omitempty"` 
    AlibabaInteractSensorAudioResponse
}

/* model for simplify = false
type AlibabaInteractSensorAudioResponse struct {

    // result=0表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorAudioResponse struct {

    // result=0表示成功
    Result   string `json:"result,omitempty"`

}
