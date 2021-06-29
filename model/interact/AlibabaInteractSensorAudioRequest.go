package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
声音 API请求
alibaba.interact.sensor.audio

客户端声音
*/
type AlibabaInteractSensorAudioRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorAudioRequest对象
func NewAlibabaInteractSensorAudioRequest() *AlibabaInteractSensorAudioRequest{
    return &AlibabaInteractSensorAudioRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorAudioRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.audio"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorAudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
