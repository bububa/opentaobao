package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
吹气 API请求
alibaba.interact.sensor.blow

客户端吹气
*/
type AlibabaInteractSensorBlowRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorBlowRequest对象
func NewAlibabaInteractSensorBlowRequest() *AlibabaInteractSensorBlowRequest{
    return &AlibabaInteractSensorBlowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorBlowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.blow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorBlowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
