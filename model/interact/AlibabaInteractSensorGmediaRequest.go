package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gmedia API请求
alibaba.interact.sensor.gmedia

媒体功能
*/
type AlibabaInteractSensorGmediaRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGmediaRequest对象
func NewAlibabaInteractSensorGmediaRequest() *AlibabaInteractSensorGmediaRequest{
    return &AlibabaInteractSensorGmediaRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGmediaRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gmedia"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGmediaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
