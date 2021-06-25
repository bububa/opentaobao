package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
takePhoto APIRequest
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
type AlibabaInteractSensorTakephotoRequest struct {
    model.Params

}

func NewAlibabaInteractSensorTakephotoRequest() *AlibabaInteractSensorTakephotoRequest{
    return &AlibabaInteractSensorTakephotoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorTakephotoRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.takephoto"
}

func (r AlibabaInteractSensorTakephotoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


