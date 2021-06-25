package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
美妆虚拟试装 APIRequest
alibaba.interact.sensor.makeup

手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
*/
type AlibabaInteractSensorMakeupRequest struct {
    model.Params

}

func NewAlibabaInteractSensorMakeupRequest() *AlibabaInteractSensorMakeupRequest{
    return &AlibabaInteractSensorMakeupRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorMakeupRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.makeup"
}

func (r AlibabaInteractSensorMakeupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


