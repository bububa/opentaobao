package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分享 API请求
alibaba.interact.sensor.share

客户端分享
*/
type AlibabaInteractSensorShareRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorShareRequest对象
func NewAlibabaInteractSensorShareRequest() *AlibabaInteractSensorShareRequest{
    return &AlibabaInteractSensorShareRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorShareRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.share"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorShareRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
