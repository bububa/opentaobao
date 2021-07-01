package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐藏titleBar API请求
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
type AlibabaInteractSensorTitlebarhideAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorTitlebarhideAPIRequest对象
func NewAlibabaInteractSensorTitlebarhideRequest() *AlibabaInteractSensorTitlebarhideAPIRequest{
    return &AlibabaInteractSensorTitlebarhideAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTitlebarhideAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.titlebarhide"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTitlebarhideAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
