package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
popwindow API请求
alibaba.interact.sensor.popwindow

popwindow
*/
type AlibabaInteractSensorPopwindowRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorPopwindowRequest对象
func NewAlibabaInteractSensorPopwindowRequest() *AlibabaInteractSensorPopwindowRequest{
    return &AlibabaInteractSensorPopwindowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorPopwindowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.popwindow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorPopwindowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
