package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摇一摇 API请求
alibaba.interact.sensor.shake

摇一摇
*/
type AlibabaInteractSensorShakeAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorShakeAPIRequest对象
func NewAlibabaInteractSensorShakeRequest() *AlibabaInteractSensorShakeAPIRequest{
    return &AlibabaInteractSensorShakeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorShakeAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.shake"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorShakeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
