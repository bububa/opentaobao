package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
canvas工具包 APIRequest
alibaba.interact.sensor.gutil

canvas工具包
*/
type AlibabaInteractSensorGutilRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGutilRequest() *AlibabaInteractSensorGutilRequest{
    return &AlibabaInteractSensorGutilRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGutilRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gutil"
}

func (r AlibabaInteractSensorGutilRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


