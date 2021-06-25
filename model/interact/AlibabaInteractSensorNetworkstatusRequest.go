package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网络状态 APIRequest
alibaba.interact.sensor.networkstatus

客户端网络状态
*/
type AlibabaInteractSensorNetworkstatusRequest struct {
    model.Params

}

func NewAlibabaInteractSensorNetworkstatusRequest() *AlibabaInteractSensorNetworkstatusRequest{
    return &AlibabaInteractSensorNetworkstatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorNetworkstatusRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.networkstatus"
}

func (r AlibabaInteractSensorNetworkstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


