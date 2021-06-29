package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端震动 API请求
alibaba.interact.sensor.vibrate

客户端震动
*/
type AlibabaInteractSensorVibrateRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorVibrateRequest对象
func NewAlibabaInteractSensorVibrateRequest() *AlibabaInteractSensorVibrateRequest{
    return &AlibabaInteractSensorVibrateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorVibrateRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.vibrate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorVibrateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
