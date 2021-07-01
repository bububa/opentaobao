package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
声音 
alibaba.interact.sensor.audio

客户端声音
*/
func AlibabaInteractSensorAudio(clt *core.SDKClient, req *interact.AlibabaInteractSensorAudioAPIRequest, session string) (*interact.AlibabaInteractSensorAudioAPIResponse, error) {
    var resp interact.AlibabaInteractSensorAudioAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
