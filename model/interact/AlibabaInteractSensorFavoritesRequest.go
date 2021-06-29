package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘开放收藏夹鉴权接口 API请求
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
*/
type AlibabaInteractSensorFavoritesRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorFavoritesRequest对象
func NewAlibabaInteractSensorFavoritesRequest() *AlibabaInteractSensorFavoritesRequest{
    return &AlibabaInteractSensorFavoritesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorFavoritesRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.favorites"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorFavoritesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
