package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端打开新页面 APIRequest
alibaba.interact.sensor.openwindow

客户端打开新页面
*/
type AlibabaInteractSensorOpenwindowRequest struct {
    model.Params

}

func NewAlibabaInteractSensorOpenwindowRequest() *AlibabaInteractSensorOpenwindowRequest{
    return &AlibabaInteractSensorOpenwindowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorOpenwindowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.openwindow"
}

func (r AlibabaInteractSensorOpenwindowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


