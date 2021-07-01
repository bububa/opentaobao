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
type AlibabaInteractSensorGmediaAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGmediaAPIRequest对象
func NewAlibabaInteractSensorGmediaRequest() *AlibabaInteractSensorGmediaAPIRequest{
    return &AlibabaInteractSensorGmediaAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gmedia"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
