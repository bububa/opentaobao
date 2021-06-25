package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
声音 APIRequest
alibaba.interact.sensor.audio

客户端声音
*/
type AlibabaInteractSensorAudioRequest struct {
    model.Params

}

func NewAlibabaInteractSensorAudioRequest() *AlibabaInteractSensorAudioRequest{
    return &AlibabaInteractSensorAudioRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorAudioRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.audio"
}

func (r AlibabaInteractSensorAudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


