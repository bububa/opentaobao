package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重力感应 API请求
alibaba.interact.sensor.gravity

native获取重力感应
*/
type AlibabaInteractSensorGravityRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGravityRequest对象
func NewAlibabaInteractSensorGravityRequest() *AlibabaInteractSensorGravityRequest{
    return &AlibabaInteractSensorGravityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGravityRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gravity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGravityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
