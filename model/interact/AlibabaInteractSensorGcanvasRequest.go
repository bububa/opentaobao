package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gcanvas API请求
alibaba.interact.sensor.gcanvas

gcanvas 功能
*/
type AlibabaInteractSensorGcanvasRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGcanvasRequest对象
func NewAlibabaInteractSensorGcanvasRequest() *AlibabaInteractSensorGcanvasRequest{
    return &AlibabaInteractSensorGcanvasRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGcanvasRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gcanvas"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGcanvasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
