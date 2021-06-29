package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆页面 API请求
alibaba.interact.sensor.login

获取登陆页面
*/
type AlibabaInteractSensorLoginRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorLoginRequest对象
func NewAlibabaInteractSensorLoginRequest() *AlibabaInteractSensorLoginRequest{
    return &AlibabaInteractSensorLoginRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorLoginRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.login"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
