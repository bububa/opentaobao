package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘开放收藏夹鉴权接口 APIRequest
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
*/
type AlibabaInteractSensorFavoritesRequest struct {
    model.Params

}

func NewAlibabaInteractSensorFavoritesRequest() *AlibabaInteractSensorFavoritesRequest{
    return &AlibabaInteractSensorFavoritesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorFavoritesRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.favorites"
}

func (r AlibabaInteractSensorFavoritesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


