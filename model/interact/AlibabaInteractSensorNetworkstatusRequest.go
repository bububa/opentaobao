package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网络状态 API请求
alibaba.interact.sensor.networkstatus

客户端网络状态
*/
type AlibabaInteractSensorNetworkstatusAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorNetworkstatusAPIRequest对象
func NewAlibabaInteractSensorNetworkstatusRequest() *AlibabaInteractSensorNetworkstatusAPIRequest{
    return &AlibabaInteractSensorNetworkstatusAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorNetworkstatusAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.networkstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorNetworkstatusAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
