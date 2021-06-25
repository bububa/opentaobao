package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIRequest
alibaba.interact.sensor.glue

视频播放
*/
type AlibabaInteractSensorGlueRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGlueRequest() *AlibabaInteractSensorGlueRequest{
    return &AlibabaInteractSensorGlueRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGlueRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.glue"
}

func (r AlibabaInteractSensorGlueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


