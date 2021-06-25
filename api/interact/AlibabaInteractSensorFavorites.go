package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
手淘开放收藏夹鉴权接口 
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
*/
func AlibabaInteractSensorFavorites(clt *core.SDKClient, req *interact.AlibabaInteractSensorFavoritesRequest, session string) (*interact.AlibabaInteractSensorFavoritesResponse, error) {
    var resp interact.AlibabaInteractSensorFavoritesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
