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
type AlibabaInteractSensorPopwindowAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorPopwindowAPIRequest对象
func NewAlibabaInteractSensorPopwindowRequest() *AlibabaInteractSensorPopwindowAPIRequest{
    return &AlibabaInteractSensorPopwindowAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.popwindow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
