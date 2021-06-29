package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端打开新页面 API请求
alibaba.interact.sensor.openwindow

客户端打开新页面
*/
type AlibabaInteractSensorOpenwindowRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorOpenwindowRequest对象
func NewAlibabaInteractSensorOpenwindowRequest() *AlibabaInteractSensorOpenwindowRequest{
    return &AlibabaInteractSensorOpenwindowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorOpenwindowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.openwindow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorOpenwindowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
