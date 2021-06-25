package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
店铺收藏 
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
*/
func AlibabaInteractShopFavor(clt *core.SDKClient, req *interact.AlibabaInteractShopFavorRequest, session string) (*interact.AlibabaInteractShopFavorResponse, error) {
    var resp interact.AlibabaInteractShopFavorAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
