package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
toast API请求
alibaba.interact.sensor.toast

toast提示
*/
type AlibabaInteractSensorToastRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorToastRequest对象
func NewAlibabaInteractSensorToastRequest() *AlibabaInteractSensorToastRequest{
    return &AlibabaInteractSensorToastRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorToastRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.toast"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorToastRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
