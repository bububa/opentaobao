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
type AlibabaInteractSensorToastAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorToastAPIRequest对象
func NewAlibabaInteractSensorToastRequest() *AlibabaInteractSensorToastAPIRequest{
    return &AlibabaInteractSensorToastAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorToastAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.toast"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorToastAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
