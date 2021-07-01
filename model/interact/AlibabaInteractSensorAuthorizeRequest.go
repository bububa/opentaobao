package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端授权页 API请求
alibaba.interact.sensor.authorize

客户端授权页
*/
type AlibabaInteractSensorAuthorizeAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorAuthorizeAPIRequest对象
func NewAlibabaInteractSensorAuthorizeRequest() *AlibabaInteractSensorAuthorizeAPIRequest{
    return &AlibabaInteractSensorAuthorizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorAuthorizeAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.authorize"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorAuthorizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
