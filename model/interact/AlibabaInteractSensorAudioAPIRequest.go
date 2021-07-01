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
type AlibabaInteractSensorAudioAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorAudioAPIRequest对象
func NewAlibabaInteractSensorAudioRequest() *AlibabaInteractSensorAudioAPIRequest{
    return &AlibabaInteractSensorAudioAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorAudioAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.audio"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorAudioAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
