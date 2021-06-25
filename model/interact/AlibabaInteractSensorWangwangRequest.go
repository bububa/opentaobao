package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘拉起旺旺接口 APIRequest
alibaba.interact.sensor.wangwang

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
*/
type AlibabaInteractSensorWangwangRequest struct {
    model.Params

}

func NewAlibabaInteractSensorWangwangRequest() *AlibabaInteractSensorWangwangRequest{
    return &AlibabaInteractSensorWangwangRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorWangwangRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.wangwang"
}

func (r AlibabaInteractSensorWangwangRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


