package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码相关API APIRequest
alibaba.interact.sensor.ma

码相关API
*/
type AlibabaInteractSensorMaRequest struct {
    model.Params

}

func NewAlibabaInteractSensorMaRequest() *AlibabaInteractSensorMaRequest{
    return &AlibabaInteractSensorMaRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorMaRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.ma"
}

func (r AlibabaInteractSensorMaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


