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
type AlibabaInteractSensorGcanvasAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGcanvasAPIRequest对象
func NewAlibabaInteractSensorGcanvasRequest() *AlibabaInteractSensorGcanvasAPIRequest{
    return &AlibabaInteractSensorGcanvasAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gcanvas"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
