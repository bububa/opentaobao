package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码相关API API请求
alibaba.interact.sensor.ma

码相关API
*/
type AlibabaInteractSensorMaRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorMaRequest对象
func NewAlibabaInteractSensorMaRequest() *AlibabaInteractSensorMaRequest{
    return &AlibabaInteractSensorMaRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorMaRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.ma"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorMaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
