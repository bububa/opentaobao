package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
陀螺仪 APIRequest
alibaba.interact.sensor.gyro

客户端陀螺仪
*/
type AlibabaInteractSensorGyroRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGyroRequest() *AlibabaInteractSensorGyroRequest{
    return &AlibabaInteractSensorGyroRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGyroRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gyro"
}

func (r AlibabaInteractSensorGyroRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


